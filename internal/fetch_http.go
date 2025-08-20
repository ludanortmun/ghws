package internal

//path = target.ToPathBase() + "/" + path
//log.Println("Will fetch:", path)
//
//res, err := http.Get(path)
//if err != nil {
//log.Printf("Error fetching %s: %v", path, err)
//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//return
//}
//
//switch res.StatusCode {
//case http.StatusOK:
//log.Printf("Successfully fetched %s", path)
//case http.StatusNotFound:
//log.Printf("Resource not found: %s", path)
//http.Error(w, "Not Found", http.StatusNotFound)
//return
//default:
//log.Printf("Unexpected status code %d for %s", res.StatusCode, path)
//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//return
//}
//
//defer res.Body.Close()
//bodyBytes, err := io.ReadAll(res.Body)
//if err != nil {
//log.Printf("Error reading response body for %s: %v", path, err)
//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//return
//}
