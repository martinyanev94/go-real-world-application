func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        userIDCookie, err := req.Cookie("identity")
        if err != nil {
            http.Error(rw, "Unauthorized", http.StatusUnauthorized)
            return
        }
        ctx := req.Context()
        ctx = ContextWithUser(ctx, userIDCookie.Value)
        req = req.WithContext(ctx)
        h.ServeHTTP(rw, req)
    })
}
