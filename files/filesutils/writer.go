package file

import "os"

func Writer(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
