package main

import (
	"fmt"
	"mastra/file-mgt/fileutil"
)

func main() {
	content, err := fileutils.ReadFile("Data/text.txt")
	fmt.Println(content)
	fmt.Println(err)	
}