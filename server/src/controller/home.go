package controller

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func HomeTemplate(w http.ResponseWriter, r *http.Request) { // later i will improve the templates for now its ready
	// i dont really want to use javascript outside of making some art or something but I will not use javascript for anything else

	file, _ := os.ReadFile("templates/index.html")
	tmpl := template.Must(template.New("a").Parse(string(file)))
	if err := tmpl.Execute(w, postAvaible); err != nil {
		fmt.Println(err)
	}
}
