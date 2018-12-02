package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goweb/pkg/setting"
)

var db *gorm.DB

type Model struct {
	gorm.Model
}

func init() {
	var (
		err                                                     error
		dbConnection, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbConnection = sec.Key("DB_CONNECTION").String()
	dbName = sec.Key("DB_DATABASE").String()
	user = sec.Key("DB_USERNAME").String()
	password = sec.Key("DB_PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("DB_TABLE_PREFIX").String()

	db, err = gorm.Open(dbConnection, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
