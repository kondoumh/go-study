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
	"strconv"
	"sync"
)

const OUT_DIR string = "_out"

type Project struct {
	Name  string `json:"projectName"`
	Count int    `json:"count"`
	Skip  int    `json:"skip"`
	Pages []Page `json:"pages"`
}

type Page struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func FetchPageList(projectName string) {
	count, err := FetchPageCount(projectName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %d\n", projectName, count)
	const LIMIT = 100
	pages := []Page{}
	for skip := 0; skip < count; skip += LIMIT {
		url := fmt.Sprintf("https://scrapbox.io/api/pages/%s?skip=%d&limit=%d&sort=updated", projectName, skip, LIMIT)
		data, err := fetchData(url)
		if err != nil {
			log.Fatal(err)
		}
		var project Project
		json.Unmarshal(data, &project)
		for _, page := range project.Pages {
			pages = append(pages, page)
		}
	}
	project := Project{}
	project.Count = count
	project.Name = projectName
	project.Skip = 0
	project.Pages = pages
	data, _ := json.Marshal(project)
	writeJson(projectName+".json", data)
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

func ReadProjectFile(projectName string) ([][]Page, error) {

	var divided [][]Page

	data, err := readFile(projectName + ".json")

	if err != nil {
		return divided, err
	}
	var project Project
	err2 := json.Unmarshal(data, &project)
	if (err2 != nil) {
		return divided, err2
	}

	chunkSize := len(project.Pages) / 3
	for i := 0; i < len(project.Pages); i += chunkSize {
		end := i + chunkSize
		if end > len(project.Pages) {
			end = len(project.Pages)
		}
		divided = append(divided, project.Pages[i:end])
	}
	return divided, nil
}

func FetchPagesWait(projectName string, pages []Page, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, page := range pages {
		fmt.Printf("%s\n", page.Title)
		FetchPageDetail(projectName, page.Title, page.Id)
	}
}

func FetchAllPages(projectName string) {
	data, err := readFile(projectName + ".json")
	if err != nil {
		log.Fatal(err)
	}
	var project Project
	json.Unmarshal(data, &project)
	index := 0
	for _, page := range project.Pages {
		fmt.Printf("%s\n", page.Title)
		FetchPageDetail(projectName, page.Title, strconv.Itoa(index))
		index++
	}
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

func FetchPageDetail(projectName string, pageName string, index string) {
	rawurl := fmt.Sprintf("https://scrapbox.io/api/pages/%s/%s", projectName, url.PathEscape(pageName))
	data, err := fetchData(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%s-%s.json", projectName, index)
	if err := writeJson(fileName, data); err != nil {
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

func readFile(fileName string) ([]byte, error) {
	raw, err := ioutil.ReadFile(OUT_DIR + "/" + fileName)
	if err != nil {
		return nil, err
	}
	return raw, err
}
