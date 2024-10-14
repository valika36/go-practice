package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/valikak/go-practice/retailer_app/Config"
	"github.com/valikak/go-practice/retailer_app/Models"
	"github.com/valikak/go-practice/retailer_app/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Order{})
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Customer{})
	r := Routes.SetupRouter()
	r.Run()
}
