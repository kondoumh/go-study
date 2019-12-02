package main

import (
	"flag"
	"github.com/kondoumh/go-study/cli/scrapbox"
)

func main() {
	limit := flag.Int("l", 10, "upper limit of fetching")
	order := flag.String("o", "updated", "ordering key {updated|views|linked}")
	skip := flag.Int("s", 0, "skip count")
	flag.Parse()
	scrapbox.FetchPages("kondoumh", limit, order, skip)
}
