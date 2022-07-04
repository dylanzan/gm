package main

import (
	"github.com/gin-gonic/gin"
	"gm/function"
	"log"
	"os"
	"os/signal"
)

func main() {

	router := gin.Default()

	function.Handler(router)

	listenAddr := "0.0.0.0:8888"

	go func() {
		router.Run(listenAddr)
	}()

	log.Printf("Listening port is: %v \n", listenAddr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("end function")

}