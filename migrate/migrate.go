package main

import (
	"github.com/AmanAmazing/theInstructor/initializers"
	"github.com/AmanAmazing/theInstructor/models"
)

func init() {
    initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
    initializers.DB.AutoMigrate(&models.User{},&models.Student{})

}
