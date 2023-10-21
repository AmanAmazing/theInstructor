package initializers

import (
	"os"

	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth


func JWTTokenAuth(){
    TokenAuth = jwtauth.New("HS256",[]byte(os.Getenv("jwtSecret")),nil)    
}
