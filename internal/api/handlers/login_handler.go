package handler

import (
	"fmt"
	"net/http"

	"github.com/Ahmeds-Library/Go-Jwt/internal/core/token"
	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/Ahmeds-Library/Go-Jwt/internal/models"
	"github.com/gin-gonic/gin"
)

var u models.User

// Login handler
// LoginHandler godoc
// @Summary      Login a user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Credentials"
// @Success      200
// @Failure      401
// @Router       /login [post]
func Login(c *gin.Context) {
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var dbPassword, dbID string

	err := database.Db.QueryRow("SELECT password, id FROM users  WHERE username = $1", u.Username).Scan(&dbPassword, &dbID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if u.Password == dbPassword {
		tokenstring, err := token.CreateToken(dbID, u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenstring})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}
}
