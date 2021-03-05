package strings

import (
	"strings"
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
