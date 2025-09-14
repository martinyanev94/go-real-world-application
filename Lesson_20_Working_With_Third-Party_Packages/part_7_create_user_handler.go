func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    db := setupDatabase()
    db.Create(&user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
