package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// HelloHandler ..
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// PingFunc ..
func PingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	group := new(errgroup.Group)

	group.Go(func() error {
		http.HandleFunc("/", HelloHandler)
		http.ListenAndServe(":8000", nil)

		for {
			sig := <- sigs
			fmt.Println()
			fmt.Println(sig)
			return errors.New("exit")
		}
	})

	group.Go(func() error {
		r := gin.Default()
		r.GET("/", PingFunc)
	
		_ = r.Run(":8090")
		return nil
	})


	if err := group.Wait(); err != nil {
		fmt.Println("Get erros: ", err)
	} else {
		fmt.Println("Exit Gracefully")
	}
}