package main

import (
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	limit := flag.Int("l", 10, "upper limit of fetching")
	flag.Parse()
	get(limit)
}

func get(limit *int) {
	fmt.Println(*limit)
	url := fmt.Sprintf("https://scrapbox.io/api/pages/kondoumh?limit=%d", *limit)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Println("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[body ]" + string(body))
}