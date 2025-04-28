package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/gin-gonic/gin"
)

// Signup handler
// SignupHandler godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Info"
// @Success      201
// @Failure      400
// @Router       /signup [post]
func Signup(c *gin.Context) {

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := database.Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)
	fmt.Println("User:", u.Username, "Password:", u.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		} else {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

