package main

import (
	"errors"
	"fmt"
)

func main() {
	myErr := errors.New("Something unexpected happend")

	fmt.Printf("Type of myErr is %T \n", myErr)
	fmt.Printf("Value of myErr is %#v \n", myErr)
}