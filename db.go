package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"

	
)

var db *gorm.DB
var err error

func connectDatabase() {
	var err error
	dataSourceName := "zabeer:Welcome123@tcp(127.0.0.1:3306)/LabDB"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Database connected.")
	//instruct gorm to look for tables matching structs in singular form
	db.SingularTable(true)
}
