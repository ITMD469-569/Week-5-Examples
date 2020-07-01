package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	
)

type Comment []struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main(){
	url := "https://jsonplaceholder.typicode.com/comments"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBodyBytes, err := ioutil.ReadAll(response.Body)

	var comments Comment
	//var comment Comment
	json.Unmarshal(responseBodyBytes, &comments)

	for i := range(comments) {
		fmt.Println(comments[i].Name)
	}




}




	
	



