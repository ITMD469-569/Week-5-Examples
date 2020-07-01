package main

import (
	"github.com/GithubUsername/week5lab1/covidApi"
)

func main() {
	countryUrl := "https://api.covid19api.com/countries"
	countries := covidApi.GetCountries(countryUrl)

	// print out country name and slug info
}
