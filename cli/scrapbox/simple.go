package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	get()
}

func get() {
	res, err := http.Get("https://scrapbox.io/api/pages/kondoumh?limit=10")
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