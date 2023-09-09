package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(templates template.Template) func(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	return func(httpWriter http.ResponseWriter, httpRequest *http.Request) {
		templates.ExecuteTemplate(httpWriter, "index.html", httpRequest)
	}
}

func main() {
	configuration := GetConfig()
	templates := template.Must(template.ParseGlob(configuration.TemplateFolder))

	http.HandleFunc("/", index(*templates))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil))
}
