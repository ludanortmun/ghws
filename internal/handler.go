package internal

import (
	"log"
	"mime"
	"net/http"
	"strings"
)

type GitHubContentFetcher interface {
	Fetch(target GitHubTarget, path string) ([]byte, error)
}

type GitHubHandler struct {
	fetcher GitHubContentFetcher
	sites   map[string]GitHubTarget
}

func NewGitHubHandler(
	fetcher GitHubContentFetcher,
) *GitHubHandler {
	return &GitHubHandler{
		fetcher: fetcher,
		sites:   make(map[string]GitHubTarget),
	}
}

func (h *GitHubHandler) AddRootSite(target GitHubTarget) *GitHubHandler {
	return h.AddSite("/", target)
}

func (h *GitHubHandler) AddSite(route string, target GitHubTarget) *GitHubHandler {
	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}
	if _, exists := h.sites[route]; exists {
		log.Printf("Warning: Overwriting existing site for route %s", route)
	}
	h.sites[route] = target

	ref := target.ref
	if ref == "" {
		ref = "default branch"
	}

	log.Printf("Added site with route %s, targetting (%s) %s/%s/%s",
		route, ref, target.owner, target.repository, target.directory)

	return h
}

func (h *GitHubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	target, path := h.getTargetAndPath(r.URL.Path)

	if target == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	path = assumeIndex(path)
	path = ensureExtension(path)
	ext := path[strings.LastIndex(path, "."):]

	log.Println("Will fetch:", path)

	bodyBytes, err := h.fetcher.Fetch(*target, path)

	if err != nil {
		if strings.Contains(err.Error(), NotFoundError) {
			log.Printf("File not found: %s in repository %s", path, target.repository)
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			log.Printf("Error fetching %s from %s: %v", path, target.repository, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	ct := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", ct)

	_, err = w.Write(bodyBytes)
	if err != nil {
		log.Printf("Error writing response for %s: %v", path, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *GitHubHandler) getTargetAndPath(originalPath string) (*GitHubTarget, string) {
	path := originalPath

	var target *GitHubTarget = nil

	for k, t := range h.sites {
		if strings.HasPrefix(path, k) {
			path = strings.TrimPrefix(path, k)
			target = &t
			break
		}
	}

	return target, path
}

// assumeIndex checks if the path is empty or root and returns the index.html path.
func assumeIndex(path string) string {
	if path == "/" || path == "" {
		path = "/index.html"
	}

	return path
}

// ensureExtension checks if the path has an extension, and if not, appends ".html".
func ensureExtension(path string) string {
	if !strings.Contains(path, ".") {
		path = path + ".html"
	}

	return path
}
