package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type Event struct {
    Title     string    `json:"title"`
    EventDate CustomTime `json:"event_date"`
}

type CustomTime struct {
    time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
    layout := `"2006-01-02"` // Define your expected date format
    t, err := time.Parse(layout, string(b))
    if err != nil {
        return err
    }
    ct.Time = t
    return nil
}

func main() {
    jsonData := []byte(`{"title": "Go Conference", "event_date": "2023-10-15"}`)

    var event Event
    err := json.Unmarshal(jsonData, &event)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    fmt.Printf("Event Title: %s\nDate: %s\n", event.Title, event.EventDate.Format("2006-01-02"))
}
