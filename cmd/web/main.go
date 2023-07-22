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

type Student struct {
	Raw       int
	Processed string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("page for GET")
		// w.Write([]byte("page for GET"))
		var tmpl_target = "index.tmpl"
		// var tmpl_target = "../templates/index.tmpl"
		// var tmpl_target = "../../templates/index.tmpl"
		// var tmpl_target = "./templates/index.tmpl"
		// var tmpl_target = "./templates/random.tmpl"
		// entries, err := os.ReadDir("./templates/")
		// entries, err := os.ReadDir("../../templates")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println("Contents")
		// for _, e := range entries {
		// 	fmt.Println(e.Name())
		// }

		// tmpl, err := template.New("something").ParseFiles(tmpl_target)
		// tmpl, err := template.New(tmpl_target).ParseFiles(tmpl_target)
		// dat, err := os.ReadFile(tmpl_target)
		tmpl, err := template.ParseFiles(tmpl_target)
		if err != nil {
			log.Println("Error parsing template: %v\n", tmpl_target)
			log.Fatalln(err)
		} else {
			// td := TemplateData{Raw: "abc", Processed: "ABC"}
			td := Student{Raw: 1, Processed: "ABC"}
			tmpl.Execute(w, td)
			// log.Println(dat)
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
				w.Write([]byte("page for POST"))
			}
		}
		defer resp.Body.Close()
	})
	http.ListenAndServe(portNumber, r)
}

// func indexHTMLTemplateVariableHandler(response http.ResponseWriter, request *http.Request) {
// 	var tmpl_target = "./templates/random.tmpl"
// 	tmplt := template.New(tmpl_target)       //create a new template with some name
// 	tmplt, _ = tmplt.ParseFiles(tmpl_target) //parse some content and generate a template, which is an internal representation
// 	p := Student{Raw: 1, Processed: "Aisha"} //define an instance with required field
// 	tmplt.Execute(response, p)               //merge template ‘t’ with content of ‘p’
// }

// func main() {
// 	fmt.Println("Starting Server for Templated response from file")
// 	http.HandleFunc("/", indexHTMLTemplateVariableHandler)
// 	http.ListenAndServe(portNumber, nil)
// }
