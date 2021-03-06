package utils

import (
	"fmt"
	"github.com/artstylecode/artcoding-go/strings"
	"github.com/jinzhu/gorm"
)

func GetDb(configPath string, connectionUrl string) *gorm.DB {
	configUtils := SysConfig{}
	configUtils.Load(configPath)
	mysqlConfig := configUtils.GetSectionConfig("mysql")

	connectUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig["user"], mysqlConfig["password"], mysqlConfig["host"], mysqlConfig["port"], mysqlConfig["dbName"])
	if !strings.IsEmptyString(connectionUrl) {
		connectUrl = connectionUrl
	}
	db, err := gorm.Open("mysql", connectUrl)
	FailOnError(err, "")

	return db
}
