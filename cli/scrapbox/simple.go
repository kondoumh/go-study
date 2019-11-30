package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	limit := flag.Int("l", 10, "upper limit of fetching")
	flag.Parse()
	get(limit)
}

func get(limit *int) {
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

	file, err := os.Create("out.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var pj bytes.Buffer
	json.Indent(&pj, []byte(body), "", " ")
	file.Write(pj.Bytes())
}
