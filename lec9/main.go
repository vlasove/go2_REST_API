package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var Users = []User{
	User{1, "bob", "1234"},
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

func GetToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "New User"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(tokenString))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", GetAll).Methods("GET")

	//ТРЕБУЕТ АУТЕНТИФИКАЦИИ
	router.Handle("/article", jwtMiddleware.Handler(http.HandlerFunc(PostArticle))).Methods("POST")

	router.HandleFunc("/auth", GetToken).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
