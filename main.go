package main

import (
	"fmt"
	"github.com/artstylecode/artcoding-go/utils"
)

func main() {
	testExcel()
}
func testExcel() {
	excelUtils := utils.ExcelUtils{}
	dateItem := excelUtils.Read("sourcesfile/test.xlsx")

	fmt.Println(dateItem)
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
