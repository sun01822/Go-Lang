package main

import (
	"fmt"
)

type User struct{
	Name string	
	Email string
}

func (u User) userDetails() string{
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func (u *User) changeEmail(newName , newEmail string){
	u.Name = newName
	u.Email = newEmail
}

func main(){
	user1 := User{Name: "John", Email: "example@gmail.com"}
	fmt.Println(user1.userDetails())
	user1.changeEmail("Sun", "example123@gmail.com")
	fmt.Println(user1.userDetails())
}
