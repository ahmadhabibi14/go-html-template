package model

import "html/template"

type Page struct {
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	Guid       string
}
