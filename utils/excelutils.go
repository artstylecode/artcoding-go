package utils

import (
	"github.com/tealeg/xlsx"
)
type ExcelUtils struct {

}

func (t *ExcelUtils) Read(name string)map[string][]string  {
	xlsxFile, err := xlsx.OpenFile(name)
	if err != nil {
		return nil
	}
	res := make(map[string][]string)
	for _,sheet := range xlsxFile.Sheets{
		rowItem := make([]string,0)
		for _,row := range sheet.Rows{
			for _,cell := range row.Cells{
				rowItem = append(rowItem,cell.Value)
			}
		}
		res[sheet.Name] = rowItem
	}
	return res
}
// Write 写入excel数据
func (t *ExcelUtils) Write(name string, list map[string][][]string)bool  {
	xlsxFile := xlsx.NewFile()
	for sheetName,items := range list{
		sheet,_ := xlsxFile.AddSheet(sheetName)
		for _,itemContents := range items{
			row := sheet.AddRow()
			for _,content := range itemContents{
				row.AddCell().SetValue(content)
			}
		}
	}

	return true
}
