package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", getConnectionString())

	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}

	db.AutoMigrate(&User{})

	db.DB()
	DB = db
}

type Validator interface {
	IsSatisfied(interface{}) bool
	DefaultMessage() string
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "db")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "root")
	pass := getParamString("db.password", "itemapp")
	dbname := getParamString("db.name", "itemapp")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", "charset=utf8mb4&parseTime=True")

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s", user, pass, protocol, host, port, dbname, dbargs)
}
