package common

import (
	"Beeper/backend/conf"
	"Beeper/backend/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	//Generate database url
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.SQL_Username,
		conf.SQL_Password,
		conf.SQL_ServerAddress,
		conf.SQL_ServerPort,
		conf.SQL_Database,
	)
	fmt.Println(args)
	//Link database
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	//Can't connect to database
	if err != nil {
		log.Panicln("Connect to database error:", err.Error())
	}
	//Auto generate table
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Panicln("AutoMigrate Table Failed:", err.Error())
	}
	DB = db
}

//Get database
func Get_db() *gorm.DB {
	return DB
}
