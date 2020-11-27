package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://github.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
		time.Sleep(time.Second)
	}

	for l := range c {
		go func(_l string) {
			time.Sleep(5 * time.Second)
			checkLink(_l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("⚠️", link, "might be down!")
		c <- link
		return
	}
	fmt.Println("✅️", link, "is up!")
	c <- link
}
