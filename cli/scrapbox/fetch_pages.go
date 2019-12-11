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

type Project struct {
	Name  string `json:"projectName"`
	Count int    `json:"count"`
}

func FetchAllPages(projectName string) {
	count, err := FetchPageCount(projectName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %d\n", projectName, count)
}

func FetchPageCount(projectName string) (int, error) {
	url := fmt.Sprintf("https://scrapbox.io/api/pages/%s?limit=1", projectName)
	data, err := fetchData(url)
	if err != nil {
		return 0, err
	}
	var project Project
	if err := json.Unmarshal(data, &project); err != nil {
		return 0, err
	}
	return project.Count, err
}

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
	data, err := fetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeJson("page_detail.json", data); err != nil {
		log.Fatal(err)
	}
}

func fetchData(rawurl string) ([]byte, error) {
	var res *http.Response
	var err error
	if name := os.Getenv("COOKIE_NAME"); name == "" {
		res, err = http.Get(rawurl)
	} else {
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
		res, err = client.Get(rawurl)
	}
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
