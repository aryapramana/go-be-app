package binrepo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-be-app/commonfile"
	"go-be-app/model"
	"net/http"
)

type BinHttpRepository interface {
	PostPhoneBinHttp(phone model.Phone) (model.BinData, error)
}

type repository struct {
	request *http.Client
}

func BinHttpObjRepository(req *http.Client) *repository {
	return &repository{req}
}

func (r *repository) PostPhoneBinHttp(phone model.Phone) (model.BinData, error) {
	var phone_entity model.Phone
	var bin_data model.BinData

	phone_entity.Brand = phone.Brand
	phone_entity.Display = phone.Display
	phone_entity.Discount = phone.Discount
	phone_entity.Rating = phone.Rating
	phone_entity.Price = phone.Price

	// map model to json
	json_body, _ := json.Marshal(phone_entity)

	// get url
	post_method := fmt.Sprintf(commonfile.GetHttpBinBaseURL() + "/post")

	// create request method
	req, req_err := http.NewRequest(http.MethodPost, post_method, bytes.NewBuffer(json_body))
	if req_err != nil {
		fmt.Println(req_err)
		return bin_data, req_err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// set request
	resp, resp_error := r.request.Do(req)
	if resp_error != nil {
		fmt.Println(resp_error)
		resp.Body.Close()
		return bin_data, resp_error
	}

	// map to model
	data := json.NewDecoder(resp.Body).Decode(&bin_data)
	if data != nil {
		fmt.Println(data)
		resp.Body.Close()
		return bin_data, data
	}

	resp.Body.Close()

	return bin_data, data

}
