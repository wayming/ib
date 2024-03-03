package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	baseUrl := "https://localhost:5000/v1/api/iserver"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response body:", err)
		return
	}

	fmt.Println(string(body))
}