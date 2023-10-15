package controllers

import (
	"net/http"

	"github.com/AmanAmazing/theInstructor/initializers"
	"github.com/AmanAmazing/theInstructor/models"
)

func PostLogin(w http.ResponseWriter,r *http.Request){
}
func PostSignup(w http.ResponseWriter,r *http.Request){
    user := models.User{}
    // checking if email already exits 
    var exists bool 
    emailCheck := initializers.DB.
        Where("email= ?",r.FormValue("email")).
        Limit(1).
        Find(&user)
    if (emailCheck.Error) || (){
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
}
