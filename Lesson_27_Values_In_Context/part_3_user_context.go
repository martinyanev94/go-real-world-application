type userKey int

const userKeyConstant userKey = 1
func ContextWithUser(ctx context.Context, user string) context.Context {
    return context.WithValue(ctx, userKeyConstant, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
    user, ok := ctx.Value(userKeyConstant).(string)
    return user, ok
}
