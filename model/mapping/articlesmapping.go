package mapping

import "time"

type ArticlesMapping struct {
	SourceNews  SourceNews `json:"source_news"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	URLToImage  string     `json:"url_image"`
	PublishedAt time.Time  `json:"published_at"`
	Content     string     `json:"content"`
}

func (articlesmapping *ArticlesMapping) SetSource(SourceNews SourceNews) {
	articlesmapping.SourceNews = SourceNews
}

func (articlesmapping ArticlesMapping) GetSource() SourceNews {
	return articlesmapping.SourceNews
}

func (articlesmapping *ArticlesMapping) SetAuthor(Author string) {
	articlesmapping.Author = Author
}

func (articlesmapping ArticlesMapping) GetAuthor() string {
	return articlesmapping.Author
}

func (articlesmapping *ArticlesMapping) SetTitle(Title string) {
	articlesmapping.Title = Title
}

func (articlesmapping ArticlesMapping) GetTitle() string {
	return articlesmapping.Title
}

func (articlesmapping *ArticlesMapping) SetDescription(Description string) {
	articlesmapping.Description = Description
}

func (articlesmapping ArticlesMapping) GetDescription() string {
	return articlesmapping.Description
}

func (articlesmapping *ArticlesMapping) SetURL(URL string) {
	articlesmapping.URL = URL
}

func (articlesmapping ArticlesMapping) GetURL() string {
	return articlesmapping.URL
}

func (articlesmapping *ArticlesMapping) SetURLToImage(URLToImage string) {
	articlesmapping.URLToImage = URLToImage
}

func (articlesmapping ArticlesMapping) GetURLToImage() string {
	return articlesmapping.URLToImage
}

func (articlesmapping *ArticlesMapping) SetPublishedAt(PublishedAt time.Time) {
	articlesmapping.PublishedAt = PublishedAt
}

func (articlesmapping ArticlesMapping) GetPublishedAt() time.Time {
	return articlesmapping.PublishedAt
}

func (articlesmapping *ArticlesMapping) SetContent(Content string) {
	articlesmapping.Content = Content
}

func (articlesmapping ArticlesMapping) GetContent() string {
	return articlesmapping.Content
}
