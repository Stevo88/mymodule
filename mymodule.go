package mypackage

import (
	"fmt"
)

func Hello() string {
	return "Hello, You!"
}

func HelloAgain() string {
	fmt.Println("and again")
	return "Hello, Again!"
}