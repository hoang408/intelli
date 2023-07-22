package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", WelcomeHandler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	Name        string
	Nationality string //unexported field.
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("welcomeform.html")
		check(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		myUser := User{}
		myUser.Name = r.Form.Get("entered_name")
		myUser.Nationality = r.Form.Get("entered_nationality")
		t, err := template.ParseFiles("welcomeresponse.html")
		check(err)
		t.Execute(w, myUser)
	}
}
