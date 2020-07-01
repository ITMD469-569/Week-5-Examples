package covidApi

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Countries []struct {
	// struct fields here
}

func makeRequest(url string) (responseBody []byte) {
	response, err := http.Get(url) // make the request
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()                             // close the connection after the function exits
	responseBodyBytes, err := ioutil.ReadAll(response.Body) // read the response
	return responseBodyBytes

}

func GetCountries(url string) (countries Countries) {
	// makes request to URL and unmarshals into Countries struct

}
