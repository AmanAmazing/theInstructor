package main

import (
	"net/http"

	"github.com/AmanAmazing/theInstructor/controllers"
	"github.com/AmanAmazing/theInstructor/initializers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm"
)


var DB *gorm.DB

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
    initializers.JWTTokenAuth()
}
func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Handle("/static/*",http.FileServer(http.Dir("static")))
    // so css filepath would be something likle href="/static/index.css" 
    r.Group(normalRoutes)
    r.Group(protectedRoutes)
    http.ListenAndServe(":3000",r)
}


func normalRoutes(r chi.Router){
    r.Get("/",func(w http.ResponseWriter,r *http.Request){
        w.Write([]byte("<h1>Have a show you stink</h1>"))
    })
    r.Post("/login",controllers.PostLogin)
    r.Get("/login",controllers.GetLogin)
    r.Post("/signup",controllers.PostSignup)
    r.Get("/signup",controllers.GetSignup)

    // experimental features
    r.Post("/expsearch",controllers.ExpPostSearch)
    r.Get("/expmovie",controllers.ExpGetMovie)

}

func protectedRoutes(r chi.Router){
    r.Use(jwtauth.Verifier(initializers.TokenAuth))
    r.Use(jwtauth.Authenticator)

    r.Get("/home",controllers.GetHome)    
    r.Get("/students",controllers.GetStudents)
}
