package endpointget

import (
	"fmt"
	"go-be-app/dbconnect"
	"go-be-app/newsrepo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPhoneData(c *gin.Context) {
	var data_error error

	phone_repository := dbconnect.NewRepository(dbconnect.InitDbConn())

	phone_all, data_error := phone_repository.FindAll()
	if data_error != nil {
		fmt.Println(data_error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   phone_all,
		"total":  len(phone_all),
	})

}

func GetDisplayPhoneData(c *gin.Context) {
	display_model := c.Query("display")
	var data_error error

	phone_repository := dbconnect.NewRepository(dbconnect.InitDbConn())

	phone_get_by_display, data_error := phone_repository.FindByDisplay(display_model)
	if data_error != nil {
		fmt.Println(data_error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"data":       phone_get_by_display,
		"total_data": len(phone_get_by_display),
	})

}

func GetNewsReport(c *gin.Context) {
	country_id := c.Query("country")
	page := c.Query("page")
	var data_error error

	news_repository := newsrepo.NewsObjRepository(dbconnect.InitApiClient())

	news_trending, http_code, data_error := news_repository.FindNewsReport(country_id, page)
	if data_error != nil {
		fmt.Println(data_error)
		return
	}

	c.JSON(http_code, gin.H{
		"status": news_trending.GetStatus(),
		"total":  len(news_trending.GetArticlesMapping()),
		"data":   news_trending.GetArticlesMapping(),
	})

}
