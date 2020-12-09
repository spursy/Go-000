package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	group := new(errgroup.Group)
	nums := []int{-1, 0 , 1}

	for _, num := range nums {
		tempNum := num

		group.Go(func() error {
			if tempNum < 0 {
				return errors.New("tempNum < 0 !!!!")
			}
			fmt.Println("tempNum: ", tempNum)
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get erros: ", err)
	} else {
		fmt.Println("Get all num successfully")
	}
}