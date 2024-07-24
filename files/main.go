package main

import (
	"fmt"
	file "mastra/filesutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/Data/"
	content, err := file.ReadFile(rootPath + "/Data/text.txt")
	if err == nil {
		fmt.Println(content)
		newContent := "Hello world. Learning go has been really fun"
		file.Writer(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error Panic!! %v", err)
	}
}
