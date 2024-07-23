package file

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)	
	if err != nil {
		fmt.Println("Failed to read File")
		return "", err
	}else {
		return string(content), nil
	}
}