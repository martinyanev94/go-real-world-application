func WithTrackingID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, "trackingID", id)
}

func TrackingIDFromContext(ctx context.Context) (string, bool) {
    return ctx.Value("trackingID").(string)
}
