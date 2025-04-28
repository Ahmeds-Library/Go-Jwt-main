package handler

import (
    "net/http"

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
// @Failure      409
// @Router       /signup [post]
func Signup(c *gin.Context) {
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := database.CreateUser(u.Username, u.Password)
    if err != nil {
        if err.Error() == "username already exists" {
            c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}