package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type User struct{
	gorm.Model
	Name string
	Age int64
	Email string
}



func main(){
	dsn := "root:@Sun2021Yes#@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("Connect to database successfully")

	// AutoMigrate the table
	err = db.AutoMigrate(&User{})
	if err != nil{
		panic("failed to migrate database")
	}
	fmt.Println("AutoMigrate the table successfully")
	

	// // Create a new user
	// singleUser:= User{Name: "Tom", Age: 18, Email: "tom@gmail.com"}
	// result:=db.Create(&singleUser)
	// if result.Error != nil{
	// 	panic("failed to create a new user")
	// }
	// fmt.Println("Row affected when creating a new user: ", result.RowsAffected)
	// fmt.Println("Auto generate ID when creating a new user: ", singleUser.ID)	

	// // // Create multiple users
	// multipleUsers:= []User{
	// 	{Name: "Shon Doe", Age: 20, Email: "johndoe@example.com"},
    // 	{Name: "Don Doe", Age: 25, Email: "johndoe@example.com"},
    // 	{Name: "Von Doe", Age: 35, Email: "johndoe@example.com"},
    // 	{Name: "Gone Doe", Age: 55,  Email: "johndoe@example.com"},
	// }

	// result:=db.Create(&multipleUsers)
	// if result.Error != nil{
	// 	panic("failed to create multiple users")
	// }
	// fmt.Println("Row affected when creating multiple users: ", result.RowsAffected)

	// // Read a user
	// var readUser User
	// result:=db.Where("name = ?", "Tom").First(&readUser)
	// // select * from users where name = 'Tom' limit 1;
	// if result.Error != nil{
	// 	panic("failed to read a user")
	// }
	// fmt.Println("Read a user successfully: ", readUser)

	// // Read all user 
	// var readUsers []User
	// result:=db.Find(&readUsers)
	// // select * from users;
	// if result.Error != nil{
	// 	panic("failed to read all users")
	// }
	// fmt.Println("Read all users successfully: ", readUsers)
	// fmt.Println("Row affected when reading all users: ", result.RowsAffected)
	

	// Update a user
	// Read a user
	// var readUser User
	// db.Where("name = ?", "NewName").First(&readUser)
	// updateUser:= User{Name: "Tom", Age: 20, Email: "tom@gmail.com"}
	// // Update users set name = 'Tom', age = 20, email = 'newEmail@gmail' where id = 1;
	// result:=db.Model(&readUser).Updates(updateUser)
	// if result.Error != nil{
	// 	panic("failed to update a user")
	// }
	// fmt.Println("Update a user successfully: ", readUser)
	// fmt.Println("Row affected when updating a user: ", result.RowsAffected)

	// Delete a user
	// Read a user
	// var readUser User
	// db.Where("name = ?", "Tom").First(&readUser)
	// // Delete from users where id = 1;
	// result:=db.Delete(&readUser)
	// if result.Error != nil{
	// 	panic("failed to delete a user")
	// }
	// fmt.Println("Delete a user successfully: ", readUser)
	// fmt.Println("Row affected when deleting a user: ", result.RowsAffected)
	
	var readAllUser []User
	db.Find(&readAllUser)
	fmt.Println("Read all users successfully: ", readAllUser)
	fmt.Println("Row affected when reading all users: ", len(readAllUser))

	
}	