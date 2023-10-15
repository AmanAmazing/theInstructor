package main

import (
	"fmt"
	"net/http"

	"github.com/AmanAmazing/theInstructor/controllers"
	"github.com/AmanAmazing/theInstructor/initializers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)


var DB *gorm.DB

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    
    r.Group(normalRoutes)
    r.Group(protectedRoutes)
    http.ListenAndServe(":3000",r)
}


func normalRoutes(r chi.Router){
    r.Post("/login",controllers.PostLogin)
    r.Get("/login",controllers.GetLogin)
    r.Post("/signup",controllers.PostSignup)
    r.Get("/signup",controllers.GetSignup)
}

func protectedRoutes(r chi.Router){

    r.Get("/secret",func (w http.ResponseWriter, r *http.Request){
        w.Write([]byte("secret page"))
        fmt.Println("testing logging")
    })     
}
