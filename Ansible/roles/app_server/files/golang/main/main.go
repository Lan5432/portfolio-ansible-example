package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
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
	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil)
	
	if errors.Is(serverErr, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if serverErr != nil {
		fmt.Printf("Error starting server, reason: %s\n", serverErr)
		os.Exit(1)
	}
}
