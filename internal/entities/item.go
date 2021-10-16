package entities

import "time"

type Item struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	CreatedAt time.Time   `json:"created_at"`
}

func NewItem(key string, value interface{}) *Item {
	return &Item{
		Key:       key,
		Value:     value,
		CreatedAt: time.Now(),
	}
}
