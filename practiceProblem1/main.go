package main

import (
	"bytes"
	"fmt"
)

type Location struct{
	latitude float64
	longitude float64
}


type Post struct{
	id int
	category string
	location Location
}

func (l *Location) setLocation(){
	l.latitude = 20.0
	l.longitude = 30.0
}

func  GenerateThumbnail(category string, data []byte) []byte {
	thumbnail := make([]byte, 0, len(data)/2)
	if category == "Photo"{
		thumbnail = append(thumbnail, data...)
	}else{
		thumbnail = append(thumbnail, bytes.Repeat([]byte{1}, 100)...)
	}
	return thumbnail
}

func main(){
	p1 := Post{id:1, category: "Technology"}
	p1.location.setLocation()
	fmt.Println(p1)
	p2 := Post{id:2, category: "Science"}
	p2.location.setLocation()
	fmt.Println(p2) 
	p3 := Post{id:3, category: "Photo"}
	p3.location.setLocation()
	fmt.Println(p3)
	thumbnails1 := GenerateThumbnail(p3.category, []byte{1,2,3,4,5})
	fmt.Println(thumbnails1)
	thumbnails2 := GenerateThumbnail(p2.category, []byte{})
	fmt.Println(thumbnails2)
}