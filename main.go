package main

import (
	"fmt"
	"net/http"
)

func  main()  {
	links := []string {
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
	}

	for _, link := range links {
		_, err := http.Get(link)

		if err != nil {
			fmt.Println("Failed to perform GET request to ", link)
		}
	}
}
