package main

import "fmt"

type UserSessionState interface {
	error
	isLoggedIn() bool
	getSessionID() int
}

type UnauthorizedError struct {
	UseID int
	SessionID int
}

func (err *UnauthorizedError) isLoggedIn() bool {
	return err.SessionID != 0
}

func (err *UnauthorizedError) getSessionID() int {
	return err.SessionID
}

func (httpErr *UnauthorizedError) Error() string {
	return fmt.Sprintf("User with id %v unauthoeized to perform this action", httpErr.UseID)
}

func validateUser(userID int) error{
	return &UnauthorizedError{userID, 18783664}
}

func main() {
	err := validateUser(1)
	if err != nil {
		fmt.Println(err)
		if errVal, ok := err.(UserSessionState); ok {
			if errVal.isLoggedIn() {
				fmt.Printf("Cleaning user session with session id %v", errVal.getSessionID())
			}
		}
	} else {
		fmt.Println("User is allowed to perform this action")
	}
}