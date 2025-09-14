// Accessing the value from context
userIDValue := ctxWithUserID.Value("userIDKey")
if uid, ok := userIDValue.(string); ok {
    fmt.Println("User ID found:", uid)
} else {
    fmt.Println("User ID not found")
}
