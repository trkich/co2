package database

import (
	"co2/common"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type MySQLConnector struct {
}

func (p *MySQLConnector) GetConnection() (db *gorm.DB, err error) {

	common.LoadEnv()

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	fmt.Println("MySQLConnector URI" + dbURI)
	return gorm.Open("mysql", dbURI)
}
