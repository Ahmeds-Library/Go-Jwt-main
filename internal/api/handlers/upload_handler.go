package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Ahmeds-Library/Go-Jwt/internal/core/token"
	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/Ahmeds-Library/Go-Jwt/internal/services"
	"github.com/gin-gonic/gin"
)

// Upload handler
// UploadHandler godoc
// @Summary      Upload and analyze a file
// @Description  Upload a file, analyze its content, and save the results to the database
// @Tags         protected
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload"
// @Success      200  {object}  map[string]interface{}  "File uploaded and analyzed successfully"
// @Failure      400  {object}  map[string]interface{}  "File not found"
// @Failure      500  {object}  map[string]interface{}  "Internal server error"
// @Router       /upload [post]
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
		return
	}
	savePath := filepath.Join("../../../uploads", file.Filename)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}

	content, err := os.ReadFile(savePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := token.DecodeToken(tokenString)
	fmt.Print(err, "Line 136")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	Username, ok1 := claims["username"].(string)
	UserID, ok2 := claims["id"].(string)

	fmt.Println(Username, UserID, "Line 149")

	if !ok1 || !ok2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	data := string(content)
	result := services.Analyze(data)

	database.SaveResult(database.Db, result, UserID, Username)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "File uploaded and analyzed successfully and also all Results are saved successfully in Database",
		"file":     file.Filename,
		"result":   result,
		"username": Username,
		"user_id":  UserID,
	})
}
