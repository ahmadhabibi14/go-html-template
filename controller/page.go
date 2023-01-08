package controller

import (
	"go-html-template/database"
	"go-html-template/model"
	"html/template"
	"log"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	thisPage := model.Page{}
	db := database.Connect()
	var rows, err = db.Query("SELECT page_title, page_content, page_date FROM pages WHERE id = 4")
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page!")
		return
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&thisPage.Title,
			&thisPage.RawContent,
			&thisPage.Date,
		)

		thisPage.Content = template.HTML(thisPage.RawContent)
	}

	parsedTemplate, _ := template.ParseFiles("page/javascript-lang/index.html")
	parsedTemplate.Execute(w, thisPage)
	defer db.Close()
}
