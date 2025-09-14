func hasField(v interface{}, fieldName string) bool {
    val := reflect.ValueOf(v)
    if val.Kind() != reflect.Struct {
        return false
    }
    return val.FieldByName(fieldName).IsValid()
}
