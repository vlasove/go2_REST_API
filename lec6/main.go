package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

//Items ...
var Items []Item = []Item{
	Item{"1", "abcd"},
}

//5. Начал реализовывать весь функционал
//GetItems ...
func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Items)
}

//GetItemID ...
func GetItemID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	find := false
	for _, item := range Items {
		if item.ID == id {
			find = true
			json.NewEncoder(w).Encode(item)
		}
	}
	if !find {
		w.WriteHeader(http.StatusNotFound) // Изменить статус код запроса на 404
	}
}

//PostItem ...
func PostItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Item
	json.Unmarshal(reqBody, &item)
	w.WriteHeader(http.StatusCreated)
	Items = append(Items, item)
}

//PutItemID ...
func PutItemID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	finded := false

	for index, item := range Items {
		if item.ID == id {
			finded = true
			reqBody, _ := ioutil.ReadAll(r.Body)
			w.WriteHeader(http.StatusAccepted)     // Изменяем статус код на 202
			json.Unmarshal(reqBody, &Items[index]) // перезаписываем всю информацию для статьи с Id
		}
	}

	if !finded {
		w.WriteHeader(http.StatusNotFound) // Изменяем статус код на 404
	}

}

//DeleteItemID ...
func DeleteItemID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	find := false

	for index, item := range Items {
		if item.ID == id {
			find = true
			w.WriteHeader(http.StatusAccepted) // Изменить статус код на 202
			Items = append(Items[:index], Items[index+1:]...)
		}
	}
	if !find {
		w.WriteHeader(http.StatusNotFound) // Изменить статус код на 404
	}

}

//2. Создал функциональные заготовки
// func GetItemID() {}

// func GetItems() {}

// func PostItem() {}

func main() {
	//1. Прописал логику своего API
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/item/{id}", GetItemID).Methods("GET")

	router.HandleFunc("/item", PostItem).Methods("POST")

	router.HandleFunc("/item/{id}", PutItemID).Methods("PUT")

	router.HandleFunc("/item/{id}", DeleteItemID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
