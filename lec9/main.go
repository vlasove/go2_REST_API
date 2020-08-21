package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var Users = []User{
	User(1, "bob", "1234"),
}

var SecretKey = []byte("secret")

type Article struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var Articles = []Article{Article{1, "First"}}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}
func PostArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", GetAll).Methods("GET")
	router.HandleFunc("/article", PostArticle).Methods("POST")
	router.HandleFunc("/auth", GetToken).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
