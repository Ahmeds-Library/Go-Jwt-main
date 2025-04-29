package handler

import (
	"net/http"

	"github.com/Ahmeds-Library/Go-Jwt/internal/core/token"
	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/Ahmeds-Library/Go-Jwt/internal/services"
	"github.com/gin-gonic/gin"
)

// GetResults handler
// GetResultsHandler godoc
// @Summary      Fetch user analysis results
// @Description  Retrieve analysis results for the authenticated user from the database
// @Tags         protected
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "User results fetched successfully"
// @Failure      401  {object}  map[string]interface{}  "Unauthorized"
// @Failure      500  {object}  map[string]interface{}  "Internal server error"
// @Router       /results [get]
func GetResults(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	username, ok1 := claims["username"].(string)
	userID, ok2 := claims["id"].(string)

	if !ok1 || !ok2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	offset, limit, page := services.PaginationHandler(c)

	results, err := database.FetchResults(username, userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch results"})
		return
	}

	total, err := database.CountResults(username, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to count results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User results fetched successfully",
		"results": results,
		"page":    page,
		"limit":   limit,
		"total":   total,
	})
}
