package scrapbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const OUT_DIR string = "_out"

func FetchPages(projectName string, limit *int, order *string, skip *int) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s?skip=%d&limit=%d&sort=%s", projectName, *skip, *limit, *order)
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
	if err := writeJson("pages.json", body); err != nil {
		log.Fatal(err)
	}
}

func FetchPageDetail(projectName string, pageName string) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s/%s", projectName, pageName)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeJson("page_detail.json", body); err != nil {
		log.Fatal(err)
	}
}

func writeJson(fileName string, data []byte) error {
	if _, err := os.Stat(OUT_DIR); os.IsNotExist(err) {
		os.Mkdir(OUT_DIR, 0777)
	}
	file, err := os.Create(OUT_DIR + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	var pj bytes.Buffer
	json.Indent(&pj, []byte(data), "", " ")
	file.Write(pj.Bytes())
	return nil
}
