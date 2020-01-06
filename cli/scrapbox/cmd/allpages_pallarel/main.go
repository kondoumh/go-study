package main

import (
	"flag"
	"fmt"
	"github.com/kondoumh/go-study/cli/scrapbox"
)

func main() {
	project := flag.String("p", "kondoumh", "project name")
	flag.Parse()
	divided, _ := scrapbox.ReadProjectFile(*project)
	for _, pages := range divided {
		fmt.Println(len(pages))
	}
}
