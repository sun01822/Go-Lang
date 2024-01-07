package main

import (
	"fmt"
)

func main(){
	// make function is used to create a slice

	sliceMake := make([]int, 3, 5)
	fmt.Println("size: ", len(sliceMake), "|| capacity: ", cap(sliceMake))

	fmt.Println("full sliceMake: ", sliceMake)

	// accessing and assigning the 1st element of the slice
	sliceMake[0] = 10

	fmt.Println("1st index value: ", sliceMake[0])
	fmt.Println("3rd index value: ", sliceMake[2])
	// error: trying to access out of size index
	//fmt.Println("5th index value: ", sliceMake[4])

	sliceMake = append(sliceMake, 1)
	fmt.Println("size: ", len(sliceMake), "|| capacity: ", cap(sliceMake))
	sliceMake = append(sliceMake, 2)
	fmt.Println("size: ", len(sliceMake), "|| capacity: ", cap(sliceMake))
	sliceMake = append(sliceMake, 3)
	fmt.Println("size: ", len(sliceMake), "|| capacity: ", cap(sliceMake))
	

	// make function is used to create a map
	mapMake := make(map[string]int)
	fmt.Println("map size(initially): ", len(mapMake))
	mapMake["one"] = 1
	mapMake["two"] = 2
	mapMake["three"] = 3
	fmt.Println("Full map: ", mapMake) 


	mapMake2 := make(map[string]int, 2)
	fmt.Println("map size(initially): ", len(mapMake2))
	// after allocating less number of elements as size of the map
	mapMake2["one"] = 1
	fmt.Println("map size(allocation less than size): ", len(mapMake2))
	mapMake2["two"] = 2

	// after allocating equal number of elements as size of the map
	fmt.Println("map size(allocation equal to size): ", len(mapMake2))
	
	// allocating more number of elements as size of the map
	mapMake2["three"] = 3
	fmt.Println("map size(allocation more than size): ", len(mapMake2))

	// make function is used to create a channel
	channelMake := make(chan int, 10)
	fmt.Println("channel size(initially): ", len(channelMake))

}