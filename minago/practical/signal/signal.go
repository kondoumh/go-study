package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer fmt.Println("done")
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, trapSignals...)
		go doMain()
		sig := <-sigCh
		fmt.Println("Got signal", sig)
}

func doMain() {
	defer fmt.Println("done infinite loop")
	for {
		time.Sleep(1 * time.Second)
	}
}