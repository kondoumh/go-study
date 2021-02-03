package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	doSomeThingParallel(3)
}

func doSomeThingParallel(workerNum int) error {
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	errCh := make(chan error, workerNum)
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if err := doSomeThingWithContext(cancelCtx, num); err != nil {
				cancel()
				errCh <- err
			}
			return
		}(i)
	}
	wg.Wait()
	close(errCh)
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

func doSomeThingWithContext(ctx context.Context, num int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	fmt.Println(num)
	return nil
}
