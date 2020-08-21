package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Items = []Item{
	Item{1, "first"},
	Item{2, "second"},
}

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var Users = []User{User{1, "bob", "1234"}}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Items)
}

var SecretKey = []byte("secret")

func GetToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "Tester"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
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

func UserRegister(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(body, &user)
	Users = append(Users, user)
	w.WriteHeader(http.StatusCreated)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(body, &user)
	for _, u := range Users {
		if u.Login == user.Login && u.Password == user.Password {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)

			claims["admin"] = true
			claims["name"] = "Tester"
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
			tokenString, err := token.SignedString(SecretKey)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(tokenString))
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("User doenst exists in database"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/items", jwtMiddleware.Handler(http.HandlerFunc(GetAll))).Methods("GET")
	router.HandleFunc("/token", GetToken).Methods("GET")
	router.HandleFunc("/register", UserRegister).Methods("POST")
	router.HandleFunc("/auth", Authenticate).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router)))
}
