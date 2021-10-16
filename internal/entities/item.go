package entities

import "time"

// Item represents the object that user cache in memory.
type Item struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	CreatedAt time.Time   `json:"created_at"`
}

// NewItem creates new Item.
func NewItem(key string, value interface{}) *Item {
	return &Item{
		Key:       key,
		Value:     value,
		CreatedAt: time.Now(),
	}
}
