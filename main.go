package main

import (
	"fmt"
	"github.com/artstylecode/artcoding-go/utils"
)

func main() {

	args := utils.GetArgsMapped("tableName", map[string]string{
		"o": "outfile",
		"p": "package",
		"t": "tableName",
	})

	for key, arg := range args {
		fmt.Printf("key:%svalue:%s \n", key, arg)
	}
}
