package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CorrectResponse struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func hitApi() (resp []CorrectResponse, err error) {

	url := "https://jsonplaceholder.cypress.io/posts"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		log.Println(err)
	}

	return
}

func testPostAPI(py CorrectResponse) (resp CorrectResponse, err error) {
	url := "https://jsonplaceholder.cypress.io/posts"
	method := "POST"

	jsonPayload, _ := json.Marshal(py)

	payload := strings.NewReader(string(jsonPayload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		log.Println(err)
	}

	return
}

func main() {

	_, err := hitApi()
	if err != nil {
		fmt.Println("invalid resp")
	}

	input := CorrectResponse{
		ID:     101,
		Title:  "recommendation",
		UserID: 12,
		Body:   "motorcycle",
	}

	result, _ := testPostAPI(input)
	if result != input {
		fmt.Println("not same")
	} else {
		fmt.Println("same")
	}

	// fmt.Println(testPostAPI())
}
