package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFileFromStrings(path string, title string, entryData []string) error {
	file, err := os.OpenFile(path+title+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range entryData {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}
