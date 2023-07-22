package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const portNumber = ":8080"

type TemplateData struct {
	Raw       string
	Processed string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("page for GET")
		var tmpl_target = "index.tmpl"
		tmpl, err := template.ParseFiles(tmpl_target)
		if err != nil {
			log.Println("Error parsing template: %v\n", tmpl_target)
			log.Fatalln(err)
		} else {
			td := TemplateData{Raw: "abc", Processed: "ABC"}
			tmpl.Execute(w, td)
			log.Println("GET response OK")
		}
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("page for POST")

		// Acquire response from target URL
		var url_dest = "https://developers.onemap.sg/commonapi/search?searchVal=revenue&returnGeom=n&getAddrDetails=n&pageNum=1"
		resp, err := http.Get(url_dest)
		if err != nil {
			log.Printf("Error querying: %v\n", url_dest)
			log.Fatalln(err)
		} else {
			fmt.Printf("Querying: %v\n", url_dest)
			// Use the html package to parse the response body from the request
			resp_body, err := ioutil.ReadAll(resp.Body)
			resp_body_str := string(resp_body)
			if err != nil {
				log.Println("Error parsing HTML response:", err)
			} else {
				log.Println(resp_body_str)
				// w.Write([]byte("page for POST"))
				var tmpl_target = "index.tmpl"
				tmpl, err := template.ParseFiles(tmpl_target)
				if err != nil {
					log.Println("Error parsing template: %v\n", tmpl_target)
					log.Fatalln(err)
				} else {
					// td := TemplateData{Raw: "123", Processed: "456"}
					td := TemplateData{Raw: resp_body_str, Processed: "456"}
					tmpl.Execute(w, td)
					log.Println("POST response OK")
				}
			}
		}
		defer resp.Body.Close()
	})
	http.ListenAndServe(portNumber, r)
}
