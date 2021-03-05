package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDb(configPath string) *gorm.DB {
	configUtils := SysConfig{}
	configUtils.Load("conf/config.ini")
	mysqlConfig := configUtils.GetSectionConfig("mysql")
	connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig["user"], mysqlConfig["password"], mysqlConfig["host"], mysqlConfig["port"], mysqlConfig["dbName"])
	db, err := gorm.Open("mysql", connectUrl)
	FailOnError(err, "")

	return db
}
