package endpointget

import (
	"fmt"
	"go-be-app/dbconnect"
	"go-be-app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UpdatePhoneData(c *gin.Context) {
	path_update := c.Param("id")

	var phone_update model.Phone
	err_message := c.ShouldBindJSON(&phone_update)

	if err_message != nil {
		var err_msg_list = []string{}
		for _, e := range err_message.(validator.ValidationErrors) {
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

	update_phone_err, err_update := phone_repository.UpdatePhone(phone_update, path_update)
	if err_update != nil {
		fmt.Println("handler: ", err_update)
		err_update_msg := fmt.Sprintf("Error in %s", err_update)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"messsage": err_update_msg,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   update_phone_err,
	})

}
