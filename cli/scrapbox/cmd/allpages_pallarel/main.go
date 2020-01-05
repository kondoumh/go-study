package main

import (
	"flag"
	"github.com/kondoumh/go-study/cli/scrapbox"
)

func main() {
	project := flag.String("p", "kondoumh", "project name")
	flag.Parse()
	scrapbox.FetchAllPages(*project)
}
