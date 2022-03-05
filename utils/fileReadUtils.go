package utils

import (
	"bufio"
	"epos_v3/domain"
	"os"
)

func ReadParseListFromFile(pathFile string, logger domain.AppLogger) []string {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.ErrorLogger.Println("[Close parseList file] Close error")
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
