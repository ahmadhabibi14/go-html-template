package controller

import (
	"go-html-template/database"
	"go-html-template/model"
	"html/template"
	"log"
	"net/http"
)

func HomePages(w http.ResponseWriter, r *http.Request) {
	var Pages = []model.Page{}
	db := database.Connect()
	var rows, err = db.Query("SELECT page_title, page_content, page_date, page_guid FROM pages")
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page!")
		return
	}

	defer rows.Close()
	for rows.Next() {
		thisPage := model.Page{}
		rows.Scan(
			&thisPage.Title,
			&thisPage.RawContent,
			&thisPage.Date,
			&thisPage.Guid,
		)

		thisPage.Content = template.HTML(thisPage.RawContent)
		Pages = append(Pages, thisPage)
	}

	parsedTemplate, _ := template.ParseFiles("index.html")
	parsedTemplate.Execute(w, Pages)
	defer db.Close()
}
