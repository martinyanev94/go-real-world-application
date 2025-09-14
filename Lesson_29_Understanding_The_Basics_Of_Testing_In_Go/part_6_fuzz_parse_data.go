func FuzzParseData(f *testing.F) {
    f.Fuzz(func(t *testing.T, input string) {
        // Your parsing function goes here
        if _, err := ParseData(strings.NewReader(input)); err != nil {
            t.Skip() // Handling inputs that result in errors
        }
    })
}
