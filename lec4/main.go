package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		w.WriteHeader(http.StatusNotFound) // Изменить статус код запроса на 404
		var erM = ErrorMessage{Message: "Not found article with that ID"}
		json.NewEncoder(w).Encode(erM)
	}
}

//PostNewArticle func for create new article
func PostNewArticle(w http.ResponseWriter, r *http.Request) {
	// {
	// 	"Id" : "3",
	// 	"Title" : "Title from json POST method",
	// 	"Author" : "Me",
	// 	"Content" : "Content from json POST method"
	// }
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article) // Считываем все из тела зпроса в подготовленный пустой объект Article
	w.WriteHeader(http.StatusCreated) // Изменить статус код запроса на 201
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article) //После добавления новой статьи возвращает добавленную
}

//DeleterArticleWithId ...
func DeleterArticleWithId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	find := false

	for index, article := range Articles {
		if article.Id == id {
			find = true
			w.WriteHeader(http.StatusAccepted) // Изменить статус код на 202
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	if !find {
		w.WriteHeader(http.StatusNotFound) // Изменить статус код на 404
		var erM = ErrorMessage{Message: "Article with that id doesn't exist"}
		json.NewEncoder(w).Encode(erM)
	}

}

//PutExistsArticle ....
func PutExistsArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idKey := vars["id"] // СТРОКА
	finded := false

	for index, article := range Articles {
		if article.Id == idKey {
			finded = true
			reqBody, _ := ioutil.ReadAll(r.Body)
			w.WriteHeader(http.StatusAccepted)        // Изменяем статус код на 202
			json.Unmarshal(reqBody, &Articles[index]) // перезаписываем всю информацию для статьи с Id
		}
	}

	if !finded {
		w.WriteHeader(http.StatusNotFound) // Изменяем статус код на 404
		var erM = ErrorMessage{Message: "Not found article with that ID. Try use POST first"}
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
	//Создадим запрос на добавление новой статьи
	myRouter.HandleFunc("/article", PostNewArticle).Methods("POST")

	//Создадим запрос на удаление статьи (гарантировано существует)
	myRouter.HandleFunc("/article/{id}", DeleterArticleWithId).Methods("DELETE")

	//ДЗ - Добавить PUT
	// Алгоритмы. Вводный курс (Т. Кормен) - это с нуля
	// Грокаем алгоритмы (А. Бхаргава)
	//
	// Кнут + Кормен (Алгоритмы. Построение и анализ) - это для ПРО
	myRouter.HandleFunc("/article/{id}", PutExistsArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
