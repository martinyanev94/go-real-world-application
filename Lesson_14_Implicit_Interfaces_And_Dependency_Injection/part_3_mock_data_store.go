type MockDataStore struct {
    data map[int]string
}

func NewMockDataStore() *MockDataStore {
    return &MockDataStore{data: make(map[int]string)}
}

func (mds *MockDataStore) Save(data string) error {
    mds.data[len(mds.data)+1] = data
    return nil
}

func (mds *MockDataStore) Get(id int) (string, error) {
    if val, ok := mds.data[id]; ok {
        return val, nil
    }
    return "", fmt.Errorf("no data found")
}

func TestDataService(t *testing.T) {
    mockStore := NewMockDataStore()
    service := NewDataService(mockStore)

    err := service.StoreData("Hello for Testing")
    if err != nil {
        t.Fatal(err)
    }

    data, err := service.RetrieveData(1)
    if err != nil || data != "Hello for Testing" {
        t.Fatalf("expected 'Hello for Testing' but got '%s'", data)
    }
}
