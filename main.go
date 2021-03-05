package main

import "github.com/artstylecode/artcoding-go/file"

func main() {

	textFileUtils := file.TextFile{}
	textFileUtils.SaveFile("test.txt", "123\r123")
}
