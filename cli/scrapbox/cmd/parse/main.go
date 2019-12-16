package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

type Project struct {
	Name  string `json:"projectName"`
	Count int    `json:"count"`
	Skip  int    `json:"skip"`
	Pages []Page `json:"pages"`
}

type Page struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Views         int    `json:"views"`
	Linked        int    `json:"linked"`
	Author        User   `json:"user"`
	Collaborators []User `json:"collaborators"`
}

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Contribute struct {
	UserId            string
	UserName          string
	PagesCreated      int
	PagesContributed  int
	ViewsCreatedPages int
	LinksCreatedPages int
}

type ContributeDetail struct {
	UserId           string
	UserName         string
	pagesCreated     []Page
	PagesContributed []Page
}

func main() {
	projectName := flag.String("p", "kondoumh", "project name")
	flag.Parse()
	pjdata, err := ioutil.ReadFile("_out/" + *projectName + ".json")
	if err != nil {
		log.Fatal(err)
	}
	var project Project
	if err := json.Unmarshal(pjdata, &project); err != nil {
		log.Fatal(err)
	}

	contribs := map[string]Contribute{}
	bar := pb.StartNew(project.Count)
	for i := 0; i < project.Count-1; i++ {
		sfx := fmt.Sprintf("-%d.json", i+1)
		bytes, err := ioutil.ReadFile("_out/" + *projectName + sfx)
		if err != nil {
			log.Fatal(err)
		}
		var page Page
		if err := json.Unmarshal(bytes, &page); err != nil {
			log.Fatal(err)
		}
		elm, contains := contribs[page.Author.Id]
		if contains {
			elm.PagesCreated++
			elm.ViewsCreatedPages += page.Views
			elm.LinksCreatedPages += page.Linked
			contribs[page.Author.Id] = elm
		} else {
			var contrib Contribute
			contrib.UserId = page.Author.Id
			contrib.UserName = page.Author.DisplayName
			contrib.PagesCreated = 1
			contrib.ViewsCreatedPages = page.Views
			contrib.LinksCreatedPages = page.Linked
			contribs[page.Author.Id] = contrib
		}
		for _, user := range page.Collaborators {
			elm, contains := contribs[user.Id]
			if contains {
				elm.PagesContributed++
				contribs[user.Id] = elm
			} else {
				var contrib Contribute
				contrib.UserId = user.Id
				contrib.UserName = user.DisplayName
				contrib.PagesContributed = 1
				contribs[user.Id] = contrib
			}
		}
		bar.Increment()
	}
	bar.Finish()
	file, err := os.Create("_out/" + *projectName + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(([]byte)("User Name,Pages Created,Pages Contributed,Views of Created Pages,Links of Created Pages\n"))
	for _, v := range contribs {
		data := fmt.Sprintf("%s,%d,%d,%d,%d\n", v.UserName, v.PagesCreated, v.PagesContributed, v.ViewsCreatedPages, v.LinksCreatedPages)
		file.Write(([]byte)(data))
	}
}
