package newsrepo

import (
	"encoding/json"
	"fmt"
	"go-be-app/commonfile"
	"go-be-app/model"
	"net/http"
)

type NewsRepository interface {
	FindNewsReport(country string) (model.NewsFile, error)
}

type repository struct {
	request *http.Client
}

func NewsObjRepository(req *http.Client) *repository {
	return &repository{req}
}

func (r *repository) FindNewsReport(country string) (model.NewsFile, error) {
	var news model.NewsFile

	// create request method
	req, req_err := http.NewRequest(http.MethodGet, commonfile.GetNewsBaseURL(), nil)
	if req_err != nil {
		fmt.Println(req_err)
		return news, req_err
	}

	// append query
	q := req.URL.Query()
	q.Add("country", country)
	q.Add("apiKey", commonfile.GetAPINewsKey())

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// assign query
	req.URL.RawQuery = q.Encode()

	// set request
	resp, resp_err := r.request.Do(req)
	if resp_err != nil {
		fmt.Println(resp_err)
		resp.Body.Close()
		return news, resp_err
	}

	data := json.NewDecoder(resp.Body).Decode(&news)
	if data != nil {
		resp.Body.Close()
		return news, data
	}

	resp.Body.Close()

	return news, data

}
