package scrapbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const OUT_DIR string = "_out"

func FetchPages(projectName string, limit *int, order *string, skip *int) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s?skip=%d&limit=%d&sort=%s", projectName, *skip, *limit, *order)
	data, err := fetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := writeJson("pages.json", data); err != nil {
		log.Fatal(err)
	}
}

func FetchPageDetail(projectName string, pageName string) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s/%s", projectName, pageName)
	data, err := fetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeJson("page_detail.json", data); err != nil {
		log.Fatal(err)
	}
}

func fetchData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
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
