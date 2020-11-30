package main

import "fmt"

type NetWorkError struct {

}

func (e *NetWorkError) Error() string {
	return "A network connection was aborted"
}

type FileSaveFailedError struct {

}

func (e *FileSaveFailedError) Error() string {
	return "The requested file could not be saved"
}

func saveFileToRemote() error {
	result := 2
	if result == 1 {
		return &NetWorkError{}
	} else if result == 2 {
		return &FileSaveFailedError{}
	} else {
		return nil
	}
}

func main() {
	switch err := saveFileToRemote(); err.(type) {
	case nil:
		fmt.Println("File successfully saved")
	case *NetWorkError:
		fmt.Println("Network Error:", err)
	case *FileSaveFailedError:
		fmt.Println("File save Error:", err)
	}
}



