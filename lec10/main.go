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

func PostToken(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	for _, u := range Users {
		if u.Login == user.Login && u.Password == user.Password {
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
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("You not in User DataBase!"))

}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	w.WriteHeader(http.StatusCreated)
	Users = append(Users, user)
	w.Write([]byte("You can /auth now!"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", GetAll).Methods("GET")

	//ТРЕБУЕТ АУТЕНТИФИКАЦИИ
	router.Handle("/article", jwtMiddleware.Handler(http.HandlerFunc(PostArticle))).Methods("POST")
	//Выбивает токен только тем кто есть в БД
	router.HandleFunc("/auth", PostToken).Methods("POST")
	//Позволяет добавиться в БД
	router.HandleFunc("/register", RegisterUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
