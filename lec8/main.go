package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//Config ...
type Config struct {
	User     string
	Password string
	DBname   string
	SSLmode  string
}

//ErrorMessage ...
type ErrorMessage struct {
	Message string
}

//Item ...
type Item struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Price  string `json:"price"`
	Title  string `json:"title"`
}

var config = Config{User: "postgres", Password: "1", DBname: "store", SSLmode: "disable"}
var connStr = fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v", config.User,
	config.Password, config.DBname, config.SSLmode)

//SelectAll ... Return slice []Item from db
func SelectAll() []Item {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Items")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var allItems []Item
	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.ID, &item.Amount, &item.Price, &item.Title)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allItems = append(allItems, item)
	}

	return allItems

}

//GetAll ... GET /items - return all items in database
func GetAll(w http.ResponseWriter, r *http.Request) {
	items := SelectAll()
	if len(items) < 1 {
		w.WriteHeader(http.StatusNoContent)
		var erM = ErrorMessage{Message: "No one items found in DB"}
		json.NewEncoder(w).Encode(erM)
	} else {
		json.NewEncoder(w).Encode(items)
	}
}

//InsertItem ... save to db new item
func InsertItem(item Item) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into Items (amount, price, title) VALUES($1, $2, $3)", item.Amount, item.Price, item.Title)
	if err != nil {
		panic(err)
	}
}

//PostItem - POST /item - add new item
func PostItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Item
	json.Unmarshal(reqBody, &item)
	InsertItem(item)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//GET /items - return all items in database
	router.HandleFunc("/items", GetAll).Methods("GET")
	//POST /item - save new item to DB
	router.HandleFunc("/item", PostItem).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
