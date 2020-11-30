package main

import "fmt"

type UnauthorizedError struct {
	UserID int
	SessionID int
	error
}

func (err *UnauthorizedError) isLoggedIn() bool {
	return err.SessionID != 0
}

func validateUser(userID int) error {
	err := fmt.Errorf("Session invalid for user id %d", userID)
	return &UnauthorizedError{userID, 18783664, err}
}

func main() {
	err := validateUser(1)
	if err != nil {
		if errVal, ok := err.(*UnauthorizedError); ok {
			fmt.Println("Is user logged in:", errVal.isLoggedIn())
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("User is allowed to perform this action!")
	}
}