type Message struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    message := Message{Title: "Greeting", Content: "Hello, World!"}
    json.NewEncoder(w).Encode(message)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/json", jsonResponseHandler)

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
