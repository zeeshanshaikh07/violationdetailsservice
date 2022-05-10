package config

import (
	"fmt"
	"log"

	"violationdetails/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {

	conf := utils.NewConfig() //Initialize config
	dbUser := conf.Database.User
	dbPass := conf.Database.Password
	dbName := conf.Database.Database
	dbHost := conf.Database.Host

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Successfully connected to database...!!!")

	//migrating defined modals to DB table
	// db.AutoMigrate(&model.TrafficViolationSystem{}, &model.ViolationDetails{})
	return db

}

func CloseDBConnection(db *gorm.DB) {
	dbMySql, err := db.DB()
	if err != nil {
		panic("Failed to close connection.")
	}

	dbMySql.Close()
}
