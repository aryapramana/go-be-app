package endpointget

import (
	"fmt"
	"go-be-app/dbconnect"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePhoneData(c *gin.Context) {
	delete_phone_id := c.Param("id")

	phone_repository := dbconnect.NewRepository(dbconnect.InitDbConn())

	delete_phone := phone_repository.DeletePhone(delete_phone_id)
	if delete_phone != nil {
		err_msg := fmt.Sprintf("Error in %s", delete_phone)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err_msg,
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})

}
