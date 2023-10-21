package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){
    var err error 
    dsn := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",os.Getenv("host"),os.Getenv("port"),os.Getenv("user"),os.Getenv("password"),os.Getenv("dbname"))
    DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{}) 
    if err != nil{
        log.Fatal("Failed to connect to database: ")
    }
}
