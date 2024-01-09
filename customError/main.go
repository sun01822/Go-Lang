package main

import (
	"fmt"
)

func main(){
	user := new(User)
	user.LastName = "Moon"
	err := myFunc(*user)
	if err==nil{
		fmt.Println("Nothing")
	}else{
		fmt.Println(err.Error())
	}
}

type MyError struct{
	Code int
	Message string
}

type User struct{
	FirstName string
	LastName string
}

func (e MyError) Error() string{
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func myFunc(user User) error{
	if user.FirstName == ""{
		return MyError{Code: 400, Message: "First name is required"}
	}
	if user.LastName == ""{
		return MyError{Code: 400, Message: "Last name is required"}
	}
	return nil
}