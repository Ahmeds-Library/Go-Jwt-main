package handler

import (
	"net/http"

	"github.com/Ahmeds-Library/Go-Jwt/internal/core/token"
	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/Ahmeds-Library/Go-Jwt/internal/models"
	"github.com/Ahmeds-Library/Go-Jwt/internal/utils"
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

	username, ok1 := claims[	"username"].(string)
	userID, ok2 := claims["id"].(string)

	if !ok1 || !ok2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Call the external PaginationHandler function
	offset, limit, page := utils.PaginationHandler(c)

	// Query with LIMIT and OFFSET for pagination
	rows, err := database.Db.Query(`
        SELECT words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels
        FROM results
        WHERE username = $1 AND user_id = $2
        LIMIT $3 OFFSET $4`, username, userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch results"})
		return
	}
	defer rows.Close()

	var results []models.Results
	for rows.Next() {
		var result models.Results
		if err := rows.Scan(
			&result.Words,
			&result.Digits,
			&result.SpecialChar,
			&result.Lines,
			&result.Spaces,
			&result.Sentences,
			&result.Punctuation,
			&result.Consonants,
			&result.Vowels); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading results"})
			return
		}
		results = append(results, result)
	}

	// Count total results for metadata
	var total int
	err = database.Db.QueryRow(`
        SELECT COUNT(*)
        FROM results
        WHERE username = $1 AND user_id = $2`, username, userID).Scan(&total)
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
