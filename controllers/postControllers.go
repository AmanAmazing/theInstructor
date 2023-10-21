package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/AmanAmazing/theInstructor/helpers"
	"github.com/AmanAmazing/theInstructor/initializers"
	"github.com/AmanAmazing/theInstructor/models"
	"gorm.io/gorm"
)

func PostLogin(w http.ResponseWriter,r *http.Request){
    user := models.User{}
    // check if password and email values are not nil  
    user.Email = r.FormValue("email")
    user.Password = r.FormValue("password")
    if user.Password == ""{
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    if user.Password == "" {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // get the user from the database
    userDb :=models.User{}
    result:= initializers.DB.Where("email",user.Email).Take(&userDb)
    if result.Error != nil{
        if errors.Is(result.Error,gorm.ErrRecordNotFound){
            // if the user is not found, redirect to signup page 
            http.Redirect(w,r,"/signup",http.StatusSeeOther)
            return
        }else{
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }
    // compare if the password is correct
    if helpers.CheckPasswordMatch(userDb.Password, user.Password) == false {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
   
    // jwt setting setting token
    _, tokenString, _ := initializers.TokenAuth.Encode(map[string]interface{}{"id":userDb.ID})
    http.SetCookie(w,&http.Cookie{
        HttpOnly: true,
        Expires: time.Now().Add(7*24*time.Hour),
        // uncomment below for https:
        // Secure:true,
        Name: "jwt",
        Value: tokenString,
    })
    http.Redirect(w,r,"/home",http.StatusSeeOther)
}


func PostSignup(w http.ResponseWriter,r *http.Request){
    var err error 
    user := models.User{}
    // parse form values 
    user.FirstName = r.FormValue("firstName")
    user.LastName = r.FormValue("lastName")
    user.Email = r.FormValue("email")
    // hashing the password
    user.Password, err = helpers.HashPassword(r.FormValue("password"))
    if err != nil{
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    // check if form values are valid 
     
    // check/post if user exists in the database
    check := initializers.DB.Where("email=?",user.Email).Take(&user)
    if check.Error != nil {
        if errors.Is(check.Error, gorm.ErrRecordNotFound){
            // the user is not in the database so we can add him in the database
            insertion :=initializers.DB.Create(&user) 
            if insertion.Error != nil{
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            http.Redirect(w,r,"/login",http.StatusSeeOther)
        }else{
            w.WriteHeader(http.StatusInternalServerError)
            return
        } 
    }
    // since the user is in the database, we will redirect them to the login page 
    w.Write([]byte("User already has an account"))
    return
}




