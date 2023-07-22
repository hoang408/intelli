package main

import (
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

type TemplateData struct {
	Raw       string
	Processed string
}

func main() {
	http.HandleFunc("/", WelcomeHandler)
	http.ListenAndServe(portNumber, nil)
}

type User struct {
	Name        string
	Nationality string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("index.tmpl")
		check(err)
		t.Execute(w, nil)
		log.Println("GET response")
	} else {
		r.ParseForm()
		td := TemplateData{}
		td.Raw = "temp1"
		td.Processed = "temp2"
		t, err := template.ParseFiles("index.tmpl")
		check(err)
		t.Execute(w, td)
		log.Println("POST response")
	}
}
