package main

import (
	"artcoding-go/utils"
	"fmt"
)

func main() {
	configUtils := utils.SysConfig{}
	configUtils.Load("conf/config.ini")
	mysqlConfig := configUtils.GetSectionConfig("prod")
	fmt.Println(mysqlConfig)
}
