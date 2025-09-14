func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Received request for:", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := mux.NewRouter()
    r.Use(loggingMiddleware)
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/about", aboutHandler)

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
