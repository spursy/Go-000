package main

import "fmt"

type HttpError struct {
	status int
	method string
}

func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something ent wrong with the %v request. Server returned %v staus.", httpError.method, httpError.status)
}

func getUserMail(userID int) (string, error) {
	return "", &HttpError{403, "GET"}
}

func main() {
	if email, err := getUserMail(1); err != nil {
		fmt.Println(err)

		if errVal := err.(*HttpError); errVal.status == 403 {
			fmt.Println("Performing session cleaning ...")
		}
	} else {
		fmt.Println("User email is ", email)
	}
}
