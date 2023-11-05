package controllers

import (
	"net/http"
	"text/template"
)



func GetSignup(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("templates/signup.html")) 
    tmpl.Execute(w, nil)
}

func GetLogin(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles("templates/login.html")) 
    tmpl.Execute(w, nil)
}

func GetHome(w http.ResponseWriter,r *http.Request){
    //tmpl := template.Must(template.ParseFiles("templates/home.html")) 
    tmpl := template.Must(template.ParseFiles(
        "templates/home.html",
        "templates/constants/navbar.html",
        "templates/constants/footer.html",
    ))
    tmpl.Execute(w, nil)
}

func GetStudents(w http.ResponseWriter,r *http.Request){
    tmpl := template.Must(template.ParseFiles(
        "templates/students.html",
        "templates/constants/navbar.html",
        "templates/constants/footer.html",
    ))
    tmpl.Execute(w, nil)
}

