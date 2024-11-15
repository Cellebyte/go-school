package main

import (
	"fmt"
	"io"
	"net/http"
)

const api = "https://themealdb.com/api/json/v1/1/random.php"

func main() {
	// TODO 1: Use http.Get to fetch the random meal from the API
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println("cannot connect to url")
		return
	}
	// TODO 2: Print the response status
	fmt.Printf("Status Code is %d\n", resp.StatusCode)
	// TODO 3: Print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("cannot read data from body")
		return
	}
	bodyString := string(body)
	fmt.Printf("Response Body: %s\n", bodyString)
}
