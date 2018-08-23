package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//Create channel to accept signal and end program
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Program started")
	//BLock the code from exiting
	<-quitChannel
	fmt.Println("Successfully exited")
}
