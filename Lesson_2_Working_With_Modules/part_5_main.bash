go mod init github.com/yourusername/yourmodule
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the home page!")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    http.Handle("/", r)
    
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}
