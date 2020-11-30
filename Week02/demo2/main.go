package main

import (
	"errors"
	"fmt"
)

func main() {
	myErr := errors.New("Something unexpected happened")
	fmt.Println(myErr)
}
