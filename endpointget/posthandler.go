package endpointget

import (
	"fmt"
	"go-be-app/binrepo"
	"go-be-app/dbconnect"
	"go-be-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostPhoneData(c *gin.Context) {
	var phone_req model.Phone
	err := c.ShouldBindJSON(&phone_req)

	// json validator
	if err != nil {
		var err_msg_list = []string{}
		for _, e := range err.(validator.ValidationErrors) {
			err_message := fmt.Sprintf("Error on field: %s | condition: %s", e.Field(), e.ActualTag())

			// put to array
			err_msg_list = append(err_msg_list, err_message)

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err_msg_list,
		})

		return

	}

	phone_repository := dbconnect.NewRepository(dbconnect.InitDbConn())

	post_phone_err := phone_repository.PostPhone(phone_req)
	if post_phone_err != nil {
		fmt.Println(post_phone_err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})

}

func PostBinPhoneData(c *gin.Context) {
	var phone_req model.Phone
	err := c.ShouldBindJSON(&phone_req)

	// json validator
	if err != nil {
		var err_msg_list = []string{}
		for _, e := range err.(validator.ValidationErrors) {
			err_message := fmt.Sprintf("Error on field: %s | condition: %s", e.Field(), e.ActualTag())

			// put to array
			err_msg_list = append(err_msg_list, err_message)

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err_msg_list,
		})

		return

	}

	phone_bin_repo := binrepo.BinHttpObjRepository(dbconnect.InitApiClient())

	phone_bin_data, phone_bin_error := phone_bin_repo.PostPhoneBinHttp(phone_req)
	if phone_bin_error != nil {
		fmt.Println(phone_bin_error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": phone_bin_data,
	})

}
