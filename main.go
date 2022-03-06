package main

import (
	"go-be-app/endpointget"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/phone/data", endpointget.GetAllPhoneData)
	v1.GET("/phone/data/list", endpointget.GetDisplayPhoneData)
	v1.GET("/api/news", endpointget.GetNewsReport)

	v1.POST("/phone/request", endpointget.PostPhoneData)
	v1.POST("/api/phone/request", endpointget.PostBinPhoneData)

	v1.PUT("/phone/request/:id", endpointget.UpdatePhoneData)

	v1.DELETE("/phone/request/:id", endpointget.DeletePhoneData)

	router.Run()
}
