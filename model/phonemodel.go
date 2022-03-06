package model

import "time"

type Phone struct {
	Id        int       `json:"id"`
	Brand     string    `json:"brand" binding:"required"`
	Display   string    `json:"display" binding:"required"`
	Price     int       `json:"price" binding:"required,number"`
	Rating    int       `json:"rating" binding:"required,number"`
	Discount  int       `json:"discount" binding:"required,number"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
