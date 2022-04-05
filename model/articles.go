package model

import "time"

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

func (articles *Articles) SetSource(Source Source) {
	articles.Source = Source
}

func (articles Articles) GetSource() Source {
	return articles.Source
}

func (articles *Articles) SetAuthor(Author string) {
	articles.Author = Author
}

func (articles Articles) GetAuthor() string {
	return articles.Author
}

func (articles *Articles) SetTitle(Title string) {
	articles.Title = Title
}

func (articles Articles) GetTitle() string {
	return articles.Title
}

func (articles Articles) SetDescription(Description string) {
	articles.Description = Description
}

func (articles Articles) GetDescription() string {
	return articles.Description
}

func (articles Articles) SetURL(URL string) {
	articles.URL = URL
}

func (articles Articles) GetURL() string {
	return articles.URL
}

func (articles *Articles) SetURLToImage(URLToImage string) {
	articles.URLToImage = URLToImage
}

func (articles Articles) GetURLToImage() string {
	return articles.URLToImage
}

func (articles *Articles) SetPublishedAt(PublishedAt time.Time) {
	articles.PublishedAt = PublishedAt
}

func (articles Articles) GetPublishedAt() time.Time {
	return articles.PublishedAt
}

func (articles *Articles) SetContent(Content string) {
	articles.Content = Content
}

func (articles Articles) GetContent() string {
	return articles.Content
}
