package strings

import (
	"strings"
	"unicode/utf8"
)

// ToHumpName 下划线转换为驼峰
func ToHumpName(str string) string {
	strList := strings.Split(str, "_")
	newstrList := make([]string, 0)
	for _, temp := range strList {
		newstrList = append(newstrList, strings.Title(temp))
	}
	return strings.Join(newstrList, "")

}

// IsEmptyString 是否为空字符串
func IsEmptyString(str string) bool {
	return utf8.RuneCountInString(str) == 0
}
