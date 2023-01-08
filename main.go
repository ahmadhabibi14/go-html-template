package main

import (
	"fmt"
	"go-html-template/controller"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server ready at http://localhost:8080/")
	http.HandleFunc("/", controller.HomePages)
	http.HandleFunc("/javascript", controller.Page)
	http.HandleFunc("/javascript/edit", controller.EditPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
