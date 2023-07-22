package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		w.Write([]byte("page for GET"))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("page for POST")
		var url_dest = "https://developers.onemap.sg/commonapi/search?searchVal=revenue&returnGeom=n&getAddrDetails=n&pageNum=1"
		resp, err := http.Get(url_dest)
		defer resp.Body.Close()
		if err != nil {
			log.Printf("Error querying: %v\n", url_dest)
			log.Fatalln(err)
		} else {
			fmt.Printf("Querying: %v\n", url_dest)
			// fmt.Println(resp)

			// Use the html package to parse the response body from the request
			// doc, err := html.Parse(resp.Body)
			resp_body, err := ioutil.ReadAll(resp.Body)
			resp_body_str := string(resp_body)
			if err != nil {
				log.Println("Error parsing HTML response:", err)
			} else {
				log.Println(resp_body_str)
				w.Write([]byte("page for POST"))
			}
		}
	})
	http.ListenAndServe(portNumber, r)
}
