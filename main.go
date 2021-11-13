package main

import (
	"log"
	"open_projects/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:12345@tcp(127.0.0.1:3306)/open_projects?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test ngesave data dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "Developer"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
