package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"

)
// Post structure
type Post struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"titile"`
	Body string `json:"body"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(response))

	var staged Post
	json.Unmarshal(response, &staged)
	// fmt.Println(staged.Body)

}