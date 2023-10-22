package main

import "fmt"

func main() {
	var a = map[string]string{"brand" : "Ford", "model" : "Mustang", "year" : "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n",a)
	fmt.Printf("b\t%v\n",b)


	//	Create maps using make() function
	var map1 = make(map[string]string)  //	The map is empty now
	map1["brand"] = "Ford"
	map1["model"] = "Mustang"
	map1["year"] = "1964"
	// map1 is no longer empty now

	map2 := make(map[string]int) //	The map is empty now
	map2["Oslo"] = 1
	map2["Bergen"] = 2
	map2["Trondheim"] = 3
	map2["Stavanger"] = 4


	fmt.Printf("map1\t%v\n",map1)
	fmt.Printf("map2\t%v\n",map2)


	// Update and add map elements
	a["year"] = "1970"	// Updating an element
	a["color"] = "red"	// Adding an element

	fmt.Println(a)


	// Remove  element from map
	delete(a, "year")
	fmt.Println(a)


}