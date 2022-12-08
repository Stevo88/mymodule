package mypackage

import (
	"flag"
	"os"
	"bufio"
)

func CheckForTestMode() []string {
	
	flag.Parse()
	fmt.Printf(flag.Arg(0))

	var inputContentAsStringArray []string
	var err error

	if (flag.Arg(0) == "test") {
		inputContentAsStringArray, err = makeArrayFromFile("input_file_sample.txt")
	} else 
	{
		inputContentAsStringArray, err = makeArrayFromFile("input_file.txt")
	}
	

	if err != nil {
		panic(err)
	}

	return inputContentAsStringArray
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