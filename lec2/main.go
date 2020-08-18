package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

//Article stuct ...
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Content string `json:"Content"`
}

type ErrorMessage struct {
	Message string `json:"Message"`
}

//Articles - local DataBase
var Articles []Article

//GET request for /articles
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hint: getAllArticles woked.....")
	json.NewEncoder(w).Encode(Articles) //ResponseWriter - место , куда пишем. Articles - кого пишем

}

//GET request for article with ID
func GetArticleWithId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	find := false
	for _, article := range Articles {
		if article.Id == vars["id"] {
			find = true
			json.NewEncoder(w).Encode(article)
		}
	}
	if !find {
		var erM = ErrorMessage{Message: "Not found article with that ID"}
		json.NewEncoder(w).Encode(erM)
	}
}

func main() {
	//Добавляю 2 статьи в свою базу
	Articles = []Article{
		Article{Id: "1", Title: "First title", Author: "First author", Content: "First content"},
		Article{Id: "2", Title: "Second title", Author: "Second author", Content: "Second content"},
	}
	fmt.Println("REST API V2.0 worked....")
	//СОздаю свой маршрутизатор на основе либы mux
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", GetAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", GetArticleWithId).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
