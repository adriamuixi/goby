package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func Home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("template/home.html")
	if err != nil {
		fmt.Fprintf(res, "Unable to load template")
	}

	user := User{Id: 1,
		Name:  "John Doe",
		Email: "example@example",
		Phone: "123456"}
	t.Execute(res, user)
}
