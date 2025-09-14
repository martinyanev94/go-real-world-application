// Creating a new context with a value
ctx := context.Background()
userID := "user123"
ctxWithUserID := context.WithValue(ctx, "userIDKey", userID)
