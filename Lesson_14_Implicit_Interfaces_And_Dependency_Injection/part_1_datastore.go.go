type DataStore interface {
    Save(data string) error
    Get(id int) (string, error)
}
type Database struct {
    // Simulate database connection
}

func (d *Database) Save(data string) error {
    // Save data to the database
    fmt.Println("Data saved to database:", data)
    return nil
}

func (d *Database) Get(id int) (string, error) {
    // Get data from the database
    return fmt.Sprintf("Data from database with id: %d", id), nil
}

type InMemoryStore struct {
    store map[int]string
}

func NewInMemoryStore() *InMemoryStore {
    return &InMemoryStore{store: make(map[int]string)}
}

func (ims *InMemoryStore) Save(data string) error {
    ims.store[len(ims.store)+1] = data
    fmt.Println("Data saved in memory:", data)
    return nil
}

func (ims *InMemoryStore) Get(id int) (string, error) {
    data, exists := ims.store[id]
    if !exists {
        return "", fmt.Errorf("no data found for id: %d", id)
    }
    return data, nil
}
