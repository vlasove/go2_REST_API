package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connStr = "user=postgres password=1 dbname=productdb sslmode=disable"

//Phone ...
type Phone struct {
	ID      int
	Model   string
	Company string
	Price   int
}

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("OK!")

	//Создадим телефон  и добавим в базу данных
	phone := Phone{Model: "IXVA -11", Company: "Sam", Price: 100}
	_, err = db.Exec("insert into Phones (model, company, price) VALUES ($1, $2, $3)", phone.Model, phone.Company, phone.Price)
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.RowsAffected())

	//Обновим объект с id = 2 - поменяем название модели
	_, err = db.Exec("update Phones set model=$1 where id=$2", "IXVA -12", 2)
	if err != nil {
		panic(err)
	}

	//Удалим телефон с id = 5
	_, err = db.Exec("delete from Phones where id=$1", 5)
	if err != nil {
		panic(err)
	}

	//Считаем из БД все
	rows, err := db.Query("select * from Phones")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var phones []Phone

	for rows.Next() {
		p := Phone{}
		err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		phones = append(phones, p)
	}
	for _, p := range phones {
		fmt.Println(p.ID, p.Model, p.Company, p.Price)
	}
}
