package mypackage

import (
	"fmt"
	"os"
	"bufio"
)

func Hello() string {
	return "Hello, You!"
}

func HelloAgain() string {
	fmt.Println("and again")
	return "Hello, Again!"
}

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