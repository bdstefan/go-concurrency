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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<- c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Failed to perform GET request on", link)
		c <- "It might be down"
		return
	}

	fmt.Println(link, "is up")
	c <- "It is up for sure"
}
