type DataService struct {
    store DataStore
}

func NewDataService(store DataStore) *DataService {
    return &DataService{store: store}
}

func (ds *DataService) StoreData(data string) error {
    return ds.store.Save(data)
}

func (ds *DataService) RetrieveData(id int) (string, error) {
    return ds.store.Get(id)
}
func main() {
    // For in-memory storage
    inMemoryStore := NewInMemoryStore()
    memoryService := NewDataService(inMemoryStore)
    memoryService.StoreData("Hello from InMemory")
    data, _ := memoryService.RetrieveData(1)
    fmt.Println(data)

    // For persistent storage
    db := &Database{}
    dbService := NewDataService(db)
    dbService.StoreData("Hello from Database")
    dbData, _ := dbService.RetrieveData(1)
    fmt.Println(dbData)
}
