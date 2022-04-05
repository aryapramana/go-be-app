package mapping

type NewsFileMap struct {
	Status          string            `json:"status"`
	TotalResults    int               `json:"total_news"`
	ArticlesMapping []ArticlesMapping `json:"articles"`
}

func (newsfilemap *NewsFileMap) SetStatus(Status string) {
	newsfilemap.Status = Status
}

func (newsfilemap NewsFileMap) GetStatus() string {
	return newsfilemap.Status
}

func (newsfilemap *NewsFileMap) SetTotalResults(TotalResults int) {
	newsfilemap.TotalResults = TotalResults
}

func (newsfilemap NewsFileMap) GetTotalResults() int {
	return newsfilemap.TotalResults
}

func (newsfilemap *NewsFileMap) SetArticlesMapping(ArticlesMapping []ArticlesMapping) {
	newsfilemap.ArticlesMapping = ArticlesMapping
}

func (newsfilemap NewsFileMap) GetArticlesMapping() []ArticlesMapping {
	return newsfilemap.ArticlesMapping
}
