package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"html/template"
)


type ExpResponseSearch struct{
    Search []ExpSearch `json:"Search"` 
    Total string `json:"totalResults"` 
    Response string `json:"Response"`
    Error string `json:"Error"`
} 
type ExpSearch struct {
    Title string `json:"Title"`
    Year string `json:"Year"`
    Runtime string `json:"Runtime"`
    Genre string `json:"Genre"`
    Director string `json:"Director"`
    Plot string `json:"Plot"`
    ImdbID string `json:"imdbID"`
    Type string `json:"Type"`
    Poster string `json:"Poster"`
    Ratings []ExpRating
}

type ExpRating struct {
    Source string `json:"Source"`
    Value string `json:"Value"`
}

func ExpPostSearch(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("templates/moviesearch.html"))
    api_key := os.Getenv("API_URL")
    searchUrl := fmt.Sprint(api_key+"&s="+r.FormValue("search"))
    fmt.Println(searchUrl) 
    res, err := http.Get(searchUrl)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    var searchResults ExpResponseSearch

    if err := json.Unmarshal(body, &searchResults); err != nil {
        log.Printf("Could not unmarshal responseBytes: %v",err)
    }

    if searchResults.Response == "False"{
        tmpl.Execute(w,nil)
    }
    tmpl.Execute(w,searchResults)
}

func ExpGetHome(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("templates/home.html")) 
    tmpl.Execute(w, nil)
}
func ExpGetMovie(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("templates/movie.html")) 
    tmpl.Execute(w, nil)
}
