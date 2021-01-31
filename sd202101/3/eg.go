package main

import (
	"golang.org/x/sync/errgroup"
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		n := i
		eg.Go(func() error {
			return do(n)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

func do(n int) error {
	if n%2 == 0 {
		return errors.New("err")
	}
	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)
	return nil
}
