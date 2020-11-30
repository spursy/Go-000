package main

import "fmt"

type UnauthorizedError struct {
	UserID int
	error
}

func validateUser(userID int) error {
	err := fmt.Errorf("Session invalid for user id %d", userID)
	return &UnauthorizedError{userID, err}
}

func main() {
	err := validateUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User is allowed to perform this action")
	}
}