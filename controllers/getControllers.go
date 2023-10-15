package controllers

import (
	"net/http"
	"text/template"
)



func GetSignup(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("assets/signup.html")) 
    tmpl.Execute(w, nil)
}

func GetLogin(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("assets/login.html")) 
    tmpl.Execute(w, nil)
}
