package model

import "time"

type NewsFile struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Articles     []Articles `json:"articles"`
}
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Articles struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}
