package model

type NewsFile struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Articles     []Articles `json:"articles"`
}

func (newsfile *NewsFile) SetStatus(Status string) {
	newsfile.Status = Status
}

func (newsfile NewsFile) GetStatus() string {
	return newsfile.Status
}

func (newsfile *NewsFile) SetTotalResults(TotalResults int) {
	newsfile.TotalResults = TotalResults
}

func (newsfile NewsFile) GetTotalResults() int {
	return newsfile.TotalResults
}

func (newsfile *NewsFile) SetArticles(Articles []Articles) {
	newsfile.Articles = Articles
}

func (newsfile NewsFile) GetArticles() []Articles {
	return newsfile.Articles
}
