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
	data := map[string][][]string{
		"test": [][]string{
			[]string{
				"姓名", "测试",
			},
		},
	}
	excelUtils.Write("test2.xlsx", data)
	fmt.Println(dateItem)
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
