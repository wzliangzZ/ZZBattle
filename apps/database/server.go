package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_ROOT, DB_PWD, DB_PORT, DB_HOST, DB_NAME, DB_TIMEOUT string
var DSN string

func init() {
	err := godotenv.Load(".db_env")
	if err != nil {
		logrus.Fatalln("Error loading .db_env file")
	}
	DB_ROOT = os.Getenv("DB_ROOT")
	DB_PWD = os.Getenv("DB_PWD")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_TIMEOUT = os.Getenv("DB_TIMEOUT")
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", DB_ROOT, DB_PWD, DB_HOST, DB_PORT, DB_NAME, DB_TIMEOUT)
	logrus.Infoln("DSN OK: ", DSN)
}

func main() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		logrus.Panic("连接数据库失败, error=" + err.Error())
		return
	}
	//延时关闭数据库连接
	fmt.Println(db)
}
