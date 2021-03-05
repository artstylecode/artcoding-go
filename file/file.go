package file

import (
	"github.com/artstylecode/artcoding-go/utils"
	"io/ioutil"
	"os"
)

// TextFile 文本文件帮助类
type TextFile struct {
}

// SaveFile 保存文件
func (t TextFile) SaveFile(filepath string, content string) bool {

	if !t.fileIsExists(filepath) {
		os.Create(filepath)
	}

	err := ioutil.WriteFile(filepath, []byte(content), os.ModePerm)
	utils.FailOnError(err, "文件写入错误")
	if err != nil {
		return false
	}

	return true
}

// fileIsExists 判断文件是否存在
func (t TextFile) fileIsExists(filepath string) bool {
	var isExist = true
	if _, err := os.Stat(filepath); os.IsExist(err) {
		isExist = false
	}
	return isExist
}
