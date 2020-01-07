package main

import (
	"flag"
	"sync"
	"github.com/kondoumh/go-study/cli/scrapbox"
)

func main() {
	project := flag.String("p", "kondoumh", "project name")
	flag.Parse()
	divided, _ := scrapbox.ReadProjectFile(*project)
	var wg sync.WaitGroup
	wg.Add(len(divided))
	for _, pages := range divided {
		go scrapbox.FetchPagesWait(*project, pages, &wg)
	}
	wg.Wait()
}
