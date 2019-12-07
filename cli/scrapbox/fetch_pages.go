package scrapbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

const OUT_DIR string = "_out"

func FetchPages(projectName string, limit *int, order *string, skip *int) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s?skip=%d&limit=%d&sort=%s", projectName, *skip, *limit, *order)
	data, err := fetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeJson("pages.json", data); err != nil {
		log.Fatal(err)
	}
}

func FetchPageDetail(projectName string, pageName string) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s/%s", projectName, pageName)
	var data []byte
	var err error
	name := os.Getenv("COOKIE_NAME")
	if name != "" {
		data, err = fetchData2(url)
	} else {
		data, err = fetchData(url)
	}
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

func fetchData2(rawurl string) ([]byte, error) {
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   os.Getenv("COOKIE_NAME"),
		Value:  os.Getenv("COOKIE_VALUE"),
		Path:   "/",
		Domain: "scrapbox.io",
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse(rawurl)
	jar.SetCookies(u, cookies)
	client := &http.Client{Jar: jar}
	res, err := client.Get(rawurl)
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
