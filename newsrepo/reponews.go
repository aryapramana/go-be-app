package newsrepo

import (
	"encoding/json"
	"fmt"
	"go-be-app/commonfile"
	"go-be-app/model"
	"go-be-app/model/mapping"
	"net/http"
)

type NewsRepository interface {
	FindNewsReport(country string, page int) (model.NewsFile, error, int)
}

type repository struct {
	request *http.Client
}

func NewsObjRepository(req *http.Client) *repository {
	return &repository{req}
}

func (r *repository) FindNewsReport(country string, page string) (mapping.NewsFileMap, int, error) {
	var news model.NewsFile
	var news_map mapping.NewsFileMap

	// create request method
	req, req_err := http.NewRequest(http.MethodGet, commonfile.GetNewsBaseURL(), nil)
	if req_err != nil {
		fmt.Println(req_err)
		return news_map, http.StatusBadRequest, req_err
	}

	// append query
	q := req.URL.Query()
	q.Add("country", country)
	q.Add("page", page)
	q.Add("apiKey", commonfile.GetAPINewsKey())

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// assign query
	req.URL.RawQuery = q.Encode()

	// set request
	resp, resp_err := r.request.Do(req)
	if resp_err != nil {
		fmt.Println("set request error: ", resp_err)
		resp.Body.Close()
		return news_map, http.StatusBadRequest, resp_err
	}

	data := json.NewDecoder(resp.Body).Decode(&news)
	if data != nil {
		fmt.Println("decode error: ", data)
		resp.Body.Close()
		return news_map, http.StatusBadRequest, data
	}

	// set to mapping
	news_map.SetStatus(news.GetStatus())
	news_map.SetTotalResults(news.GetTotalResults())

	var article_list []model.Articles
	article_list = append(article_list, news.GetArticles()...)

	var article_list_mapper []mapping.ArticlesMapping

	for _, list := range article_list {
		var article_mapper mapping.ArticlesMapping
		var source_mapper mapping.SourceNews

		article_mapper.SetAuthor(list.GetAuthor())
		article_mapper.SetContent(list.GetContent())
		article_mapper.SetDescription(list.GetDescription())
		article_mapper.SetTitle(list.GetTitle())
		article_mapper.SetURL(list.GetURL())
		article_mapper.SetPublishedAt(list.GetPublishedAt())
		article_mapper.SetURLToImage(list.GetURLToImage())

		// set source
		source_mapper.SetId(list.GetSource().GetID())
		source_mapper.SetName(list.GetSource().GetName())

		// set to list
		article_mapper.SetSource(source_mapper)
		article_list_mapper = append(article_list_mapper, article_mapper)

	}

	news_map.SetArticlesMapping(article_list_mapper)

	resp.Body.Close()

	return news_map, http.StatusOK, data

}
