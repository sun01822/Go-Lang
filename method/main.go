package main

import (
	"fmt"
)

type User struct{
	Name string
	Email string
}

type Admin struct{
	Name string
	Email string
}

func (u User) userDetails() string{
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

func (a Admin) userDetails() string{
	return fmt.Sprintf("Name: %s, Email: %s", a.Name, a.Email)
}

func main(){	
	user := User{Email: "example@gmail.com", Name: "Sun"}
	fmt.Println(user.userDetails())
	admin := Admin{Name: "Admin", Email: "admin@gmail.com"}
	fmt.Println(admin.userDetails())
}
