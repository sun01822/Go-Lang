package main

import (
	"fmt"
)

type UserDetails struct {
	Name	string
	Email 	string
}


func (u UserDetails) userDetails() string {
	return fmt.Sprintf("User name is: %s and Email is: %s", u.Name, u.Email)
}


func main() {
	user1 := UserDetails{Name: "Sun", Email: "example@gmail.com"}
	fmt.Println(user1.userDetails())
}