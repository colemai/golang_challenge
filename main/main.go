/*
Purpose: Call and test client library

As will soon become apparent, I have no experience writing Go
Or any other compiled language
Let's see how we go



*/

package main

import (
	"fmt"
	// "log"
	"net/http"
	"io/ioutil"
	// "os"
	// "src/client"
)

const baseURL = "https://api.form3.tech/"

func main() {
	var response string = fetchReq()
	fmt.Println(response)
}

func fetchReq () string {
	var output string
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=EUR")
	if err != nil {
		fmt.Printf("Fetch Request failed with error: %s \n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			output = string(data)
			// fmt.Println(string(data))
		}
	return output
}