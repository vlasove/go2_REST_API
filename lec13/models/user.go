package models

import (
	"lec13/config"

	_ "github.com/lib/pq" //engine
)

//GetAllUsers ...
func GetAllUsers(users *[]User) error {
	if err := config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ...
func CreateUser(user *User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ...
func GetUserByID(user *User, id string) error {
	if err := config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ...
func UpdateUser(user *User, id string) error {
	if err := config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

//DeleteUser ...
func DeleteUser(user *User, id string) error {
	if err := config.DB.Where("id=?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
