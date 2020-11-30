package main

import "fmt"

type MyError struct {

}

func (myErr *MyError) Error() string {
	return "Something unexpected happend!"
}

func main() {
	myErr := &MyError{}

	fmt.Println(myErr)
}