package main

import "flag"

import (
	"github.com/kondoumh/go-study/cli/scrapbox"
)

func main() {
	project := flag.String("r", "kondoumh", "project name")
	page := flag.String("p", "Dev", "page title")
	flag.Parse()

	scrapbox.FetchPageDetail(*project, *page, 0)
}