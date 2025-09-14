func TestRemoteSolver(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "5.0") // Mocking a response from the server
    }))
    defer server.Close()

    rs := RemoteSolver{
        MathServerURL: server.URL,
        Client:        server.Client(),
    }

    result, err := rs.Resolve(context.Background(), "2+3")
    if err != nil {
        t.Fatal(err)
    }

    if result != 5.0 {
        t.Errorf("expected %f, got %f", 5.0, result)
    }
}
