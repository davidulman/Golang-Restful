package controllers

import (
	"net/http"

	"github.com/davidulman/Golang-Restful/initializers"
	"github.com/davidulman/Golang-Restful/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(c *gin.Context) {
	var body struct {
		FirstName string 
		LastName  string  
		Email     string  
		Password  string 
	}

	// Bind the request body to the `body` struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
	}

	res := initializers.DB.Create(&user)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
