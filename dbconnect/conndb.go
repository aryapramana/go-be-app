package dbconnect

import (
	"fmt"
	"go-be-app/model"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDbConn() *gorm.DB {
	// insert your db conn here
	dsn := "{{connection_host}}"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB error!")
	}

	fmt.Println("Ready to go!")

	db.AutoMigrate(&model.Phone{})

	return db

}

func InitApiClient() *http.Client {
	return &http.Client{}
}
