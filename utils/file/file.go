package file

import (
	"bufio"
	"os"
)

func GetLines(filePath string) (lines []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, each_ln := range text {
		lines = append(lines, each_ln)
	}

	return lines, nil
}
