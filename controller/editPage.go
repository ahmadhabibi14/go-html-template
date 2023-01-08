package controller

import (
	"html/template"
	"net/http"
)

func EditPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("page/javascript-lang/edit.html")
	parsedTemplate.Execute(w, nil)
}
