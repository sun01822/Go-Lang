package main

import (
	"fmt"
)

type Location struct{
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type Post struct{
	Category string `json:"category"`
	Time int64 `json:"time"`
	L Location `json:"location"` 
}

/* if we don't change the method receiver
	as a pointer it doesn't change the value of Post
	we need to use
	func (post *Post) setLocation(){
		post.L = Location{
		Latitude: 23.99,
		Longitude: 90.45,
	}
*/
func (post Post) setLocation(){
	post.L = Location{
		Latitude: 23.99,
		Longitude: 90.45,
	}
}

func main(){
	post:= Post{}
	post.setLocation()
	fmt.Println(post)
	fmt.Println(post.L)
}