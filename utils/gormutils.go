package utils

import "github.com/jinzhu/gorm"

func GetDb(configPath string) *gorm.DB {
	db, err := gorm.Open("mysql", "erp_beta_yuanben:tJPd7WFYMXmab7LX@tcp(main:3306)/erp_beta_yuanben?charset=utf8mb4&parseTime=True&loc=Local")
	FailOnError(err, "")

	return db
}
