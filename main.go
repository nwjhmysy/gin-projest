package main

import (
	"gin-project/app/boot"
	"log"
	"os"
	"os/signal"
)

func main() {

	boot.GinServer()

	WaitExit()
}

func WaitExit() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server.")
}
