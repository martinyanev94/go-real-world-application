func handler(rw http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    user, ok := UserFromContext(ctx)
    if !ok {
        http.Error(rw, "Unauthorized", http.StatusUnauthorized)
        return
    }

    data := req.URL.Query().Get("data")
    result, err := BusinessLogic(ctx, user, data)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    rw.Write([]byte(result))
}
