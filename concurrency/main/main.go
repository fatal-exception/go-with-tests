package main

import (
	"fmt"
	"net/http"

	"github.com/fatal-exception/go-with-tests/concurrency"
)

func main() {
	websiteIs200 := func(s string) bool {
		resp, err := http.Get(s)
		result := err == nil && resp.StatusCode == 200
		fmt.Printf("Checking %s... %v!\n", s, result)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		return result
	}
	websitesToCheck := []string{
		"https://www.google.com",
		"https://www.scabbooballoo.beans",
	}
	concurrency.CheckWebsites(websiteIs200, websitesToCheck)
}
