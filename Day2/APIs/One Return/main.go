package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/comments/1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBodyBytes, err := ioutil.ReadAll(response.Body)

	var comment Comment
	json.Unmarshal(responseBodyBytes, &comment)
	fmt.Println(comment)
}
