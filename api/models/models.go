package models

import "time"

type Contact struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone_number"`
	Subject   string    `json:"subject"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type Destination struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Ratings  float32 `json:"ratings"`
	Category string  `json:"category"`
	Pricing  float64 `json:"pricing"`
}
