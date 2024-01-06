package main

import (
	"fmt"
	"time"
)

func ping(c chan string){
	links := []string{"https://www.google.com",
	 "https://www.facebook.com", 
	 "https://www.youtube.com", 
	 "https://www.amazon.com", 
	 "https://www.wikipedia.org", 
	 "https://www.reddit.com", 
	 "https://www.yahoo.com", 
	 "https://www.twitter.com", 
	 "https://www.instagram.com", 
	 "https://www.linkedin.com"}

	 for _, link := range links {
		time.Sleep(1 * time.Second)
		c <- link
	 }
}

func main() {
	fmt.Println("Hello, World!")
	linkChannel := make(chan string, 5)
	go ping(linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println(<-linkChannel)
	fmt.Println("Main function ends")
}