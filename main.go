package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
    if err!=nil{
        log.Fatalf("loading env error!: %s",err)
    }

}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    
    r.Group(normalRoutes)
    r.Group(protectedRoutes)
    http.ListenAndServe(":3000",r)
}


func normalRoutes(r chi.Router){
    
    r.Get("/home",func(w http.ResponseWriter,r *http.Request){
        w.Write([]byte("home page"))
    })
}

func protectedRoutes(r chi.Router){
    r.Get("/secret",func (w http.ResponseWriter, r *http.Request){
        w.Write([]byte("secret page"))
        fmt.Println("testing logging")
    })     
}
