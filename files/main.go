package main

import (
	"fmt"
	file "mastra/filesutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	content, err := file.ReadFile(rootPath + "/Data/text.txt")
	if err == nil {
		fmt.Println(content)
	} else {
		fmt.Printf("Error Panic!! %v", err)
	}
}
