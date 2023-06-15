// Package main - goMonitoring
// Author: Kuntal Bar
// Description: This file is part of the goMonitoring project, which monitors Docker resources
//              and provides a web-based interface for monitoring Linux servers.

package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Body string
}



func main() {	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData {
			Title: "Monitoring",
			Body: "Test",
		}

		renderTemplate(w, "templates/index.html", data)
	})

	http.ListenAndServe(":8080", nil)
}


func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
