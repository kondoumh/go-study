package main

import (
	"flag"
	"sync"
	"time"
	"log"

	"github.com/kondoumh/go-study/cli/scrapbox"
)

func main() {
	project := flag.String("p", "kondoumh", "project name")
	flag.Parse()
	divided, _ := scrapbox.SplitProjectPages(*project, 3)
	var wg sync.WaitGroup

	start := time.Now()
	wg.Add(len(divided))
	for _, pages := range divided {
		go scrapbox.FetchPagesWait(*project, pages, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
