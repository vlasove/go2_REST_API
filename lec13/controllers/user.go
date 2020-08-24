package controllers

import (
	"fmt"
	"lec13/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers ...
func GetUsers(c *gin.Context) {
	var users []models.User
	err := models.GetAllUsers(&users)
	if err != nil {
		// log.Fatal(err)
		// panic(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

//CreateUser ...
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user) // ioutil.ReadAll(r.Body) + json.Unmarshal(...)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

//GetUserByID ...
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ...
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

//DeleteUser ...
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusAccepted, gin.H{"id " + id: "deleted"})
	}
}
