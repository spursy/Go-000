package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	originError := errors.New("I am original error message")
	err := errors.Wrap(originError, "Context Info Using Wrap")
	fmt.Printf("normal => %v \n\n", err)

	fmt.Printf("With stack trace => %+v \n\n", err)

	originalErrorUnWrapped := errors.Cause(err)
	fmt.Println(originalErrorUnWrapped)
}