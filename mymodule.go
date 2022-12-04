package mypackage

import (
	"os"
	"bufio"
)

func makeArrayFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordArray []string
	for scanner.Scan() {
		wordArray = append(wordArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return wordArray, err
}