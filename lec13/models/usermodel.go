package models

//User ...
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

//TableName ...
func (b *User) TableName() string {
	return "Users"
}

// CREATE TABLE Users (
// 	id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
// 	name varchar (50) NOT NULL,
// 	email varchar (50) NOT NULL,
// 	phone varchar (50) NOT NULL,
// 	country varchar (50) NOT NULL
// 	);
