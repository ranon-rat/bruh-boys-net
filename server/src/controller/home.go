package controller

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func HomeTemplate(w http.ResponseWriter, r *http.Request) { // later i will improve the templates for now its ready
	// i dont really want to use javascript outside of making some art or something but I will not use javascript for anything else
	agent := r.Header.Get("User-Agent")
	if strings.Contains(agent, "Mobile") {
		fmt.Fprintf(w, "<html><body><h1>no phone fags allowed</h1></body></html>")
		return
	}
	file, _ := os.ReadFile("templates/index.html")
	tmpl := template.Must(template.New("a").Parse(string(file)))
	if err := tmpl.Execute(w, postAvaible); err != nil {
		fmt.Println(err)
	}
}
