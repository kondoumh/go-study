package main

import (
	"os"
	"github.com/tenntenn/gohandson/greeting"
)

func main() {
	var g greeting.Greeting
	g.Do(os.Stdout)
}