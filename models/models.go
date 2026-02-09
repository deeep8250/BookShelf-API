package models

import "time"

type Books struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
