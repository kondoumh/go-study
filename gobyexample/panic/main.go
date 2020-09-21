package main

import "os"

func main() {
	panic("a problemm")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
