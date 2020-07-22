package main

import "fmt"

func main(){

	countries := map[string]string{
		"IN":"New Delhi",
		"US": "Washington DC",
	}

	for country, capital := range countries{
		fmt.Printf("Capital of %s is %s",country, capital)
		fmt.Println()
	}


	region_capitals := map[string]map[string]string{
		"IN": map[string]string{ 
			"KA": "Bangalore", 
			"TN" : "Chennai",
		},
		"US": map[string]string{ 
			"AR": "Memphis", 
			"CA" : "San Francisco",
		},
	}


}