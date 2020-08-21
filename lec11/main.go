package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

//Article ...
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:'title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

//User ....
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

//Config ...
type Config struct {
	Port      string
	User      string
	Password  string
	DBname    string
	SSLmode   string
	SecretKey []byte
}

var config = Config{
	Port:      ":8080",
	User:      "postgres",
	Password:  "1",
	DBname:    "articledb",
	SSLmode:   "disable",
	SecretKey: []byte("ultrarestapiv3.0"),
}

//ConnStr ...
var ConnStr = fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
	config.User, config.Password, config.DBname, config.SSLmode)

//InsertUser ...
func InsertUser(user User) {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into Users (login, password) VALUES ($1, $2)",
		user.Login, user.Password)

	if err != nil {
		log.Fatal(err)
	}

}

//UserRegister ...
func UserRegister(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	InsertUser(user)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration OK!"))
}

//SelectAll ...
func SelectAll() []User {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users
}

//PostToken ...
func PostToken(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	users := SelectAll() // SelectUserByLogin(user.Login)
	for _, u := range users {
		if u.Login == user.Login && u.Password == user.Password {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)

			claims["admin"] = true
			claims["name"] = user.Login
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
			tokenString, err := token.SignedString(config.SecretKey)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(tokenString))
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("User doesn't exists in database"))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

//SelectAllArticles ...
func SelectAllArticles() []Article {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Articles")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	articles := []Article{}

	for rows.Next() {
		a := Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			fmt.Println(err)
			continue
		}
		articles = append(articles, a)
	}

	return articles

}

//GetAll ...
func GetAll(w http.ResponseWriter, r *http.Request) {
	articles := SelectAllArticles()
	json.NewEncoder(w).Encode(articles)
}

//InsertArticle  ...
func InsertArticle(article Article) {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into Articles (title, author, content) VALUES ($1, $2, $3)",
		article.Title, article.Author, article.Content)
	if err != nil {
		log.Fatal(err)
	}
}

//PostArticle ...
func PostArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	InsertArticle(article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

//DeleteArticleFromDB ...
func DeleteArticleFromDB(id int) {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("delete from Articles where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteArticle ...
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	DeleteArticleFromDB(id)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//POST /register
	router.HandleFunc("/register", UserRegister).Methods("POST")
	//POST /auth
	router.HandleFunc("/auth", PostToken).Methods("POST")

	//GET /articles
	router.HandleFunc("/articles", GetAll).Methods("GET")

	//POST /article + json
	router.Handle("/article", jwtMiddleware.Handler(http.HandlerFunc(PostArticle))).Methods("POST")

	//DELETE /article/{id}
	router.Handle("/article/{id}", jwtMiddleware.Handler(http.HandlerFunc(DeleteArticle))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(config.Port, router))
}
