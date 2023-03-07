package common

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	var lines [][]byte

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			lines = append(lines, []byte(line))
		}
		
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	return lines

}