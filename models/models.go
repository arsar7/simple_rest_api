package models

import (
	"errors"
	"net/http"
	"strconv"

	"example/simple_rest_api/database"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, database.Users)

}

func AddUser(context *gin.Context) {
	var newUser database.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	database.Users = append(database.Users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)

}

func GetUser(context *gin.Context) {

	id, e := strconv.Atoi(context.Param("id"))
	if e != nil {
		panic(e)
	}

	user, err := getUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)

}

func getUserById(id int) (*database.User, error) {

	for _, t := range database.Users {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, errors.New("Could not find item")
}

func DeleteUser(context *gin.Context) {
	id, e := strconv.Atoi(context.Param("id"))
	if e != nil {
		panic(e)
	}

	user, err := getUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	user.IsActive = false

	context.IndentedJSON(http.StatusOK, user)
}

func Index(context *gin.Context) {

	msg := "Landing Page"

	context.IndentedJSON(http.StatusOK, msg)
}
