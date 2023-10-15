package initializers

import (
	"log"

	"github.com/joho/godotenv"
)


func LoadEnvVariables(){
	err := godotenv.Load(".env")
    if err!=nil{
        log.Fatalf("loading env error!: %s",err)
    }
}
