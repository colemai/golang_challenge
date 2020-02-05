/*
Purpose: Client Library for the specified API

*/


package library

import (
	"fmt"
	// "log"
	"net/http"
	"io/ioutil"
	// "os"
)


// Create Request
func create_req () {
	response, err = http.Get("https://api.coinbase.com/v2/prices/spot?currency=EUR")
	if err != nil {
		fmt.Printf("Create Request failed with error: %s \n", err)
		} else {
			data, err = ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
}


// Fetch Request

// List Request

// Delete Request