package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

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

func main() {
	bytes, err := ioutil.ReadFile("_out/page_detail.json")
	if err != nil {
		log.Fatal(err)
	}
	var page Page
	if err := json.Unmarshal(bytes, &page); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s : %s\n", page.Id, page.Title)
	fmt.Printf("%d %d\n", page.Views, page.Linked)
	fmt.Printf("Author : %s %s %s\n", page.Author.Id, page.Author.Name, page.Author.DisplayName)
	fmt.Printf("Collaborators:\n")
	for _, user := range page.Collaborators {
		fmt.Printf("%s %s %s\n", user.Id, user.Name, user.DisplayName)
	}
}
