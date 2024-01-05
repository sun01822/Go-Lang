package main

import (
	"bytes"
	"fmt"
)

type Location struct {
	latitude  float64
	longitude float64
}

type ThumbnailGenerator interface {
	GenerateThumbnail(data []byte) []byte
}

type PhotoThumbnailGenerator struct{}

func (ptg PhotoThumbnailGenerator) GenerateThumbnail(data []byte) []byte {
	// Implementation for generating photo thumbnails
	return append([]byte{}, data...)
}

type DefaultThumbnailGenerator struct{}

func (dtg DefaultThumbnailGenerator) GenerateThumbnail(data []byte) []byte {
	// Implementation for generating default thumbnails
	return append([]byte{}, bytes.Repeat([]byte{1}, 100)...)
}

type Post struct {
	id         int
	category   string
	location   Location
	ThumbnailGenerator
}

func (l *Location) setLocation() {
	l.latitude = 20.0
	l.longitude = 30.0
}

func NewPost(id int, category string, generator ThumbnailGenerator) Post {
	return Post{id: id, category: category, ThumbnailGenerator: generator}
}

func main() {
	p1 := NewPost(1, "Technology", DefaultThumbnailGenerator{})
	p1.location.setLocation()
	fmt.Println(p1)

	p2 := NewPost(2, "Science", DefaultThumbnailGenerator{})
	p2.location.setLocation()
	fmt.Println(p2)

	p3 := NewPost(3, "Photo", PhotoThumbnailGenerator{})
	p3.location.setLocation()
	fmt.Println(p3)

	thumbnails1 := p3.GenerateThumbnail([]byte{1, 2, 3, 4, 5})
	fmt.Println(thumbnails1)

	thumbnails2 := p2.GenerateThumbnail([]byte{})
	fmt.Println(thumbnails2)
}
