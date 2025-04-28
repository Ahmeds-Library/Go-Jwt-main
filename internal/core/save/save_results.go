package save_results

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/Ahmeds-Library/Go-Jwt/internal/models"
)

func SaveResult(db *sql.DB, result models.Results, userID string, username string) error {

	query := database.SaveResultQuery

	_, err := db.Exec(query,
		result.Words,
		result.Digits,
		result.SpecialChar,
		result.Lines,
		result.Spaces,
		result.Sentences,
		result.Punctuation,
		result.Consonants,
		result.Vowels,
		userID,
		username,
	)

	if err != nil {
		log.Println("Error inserting result:", err)
		return err
	}

	fmt.Println("Result saved successfully.")
	return nil
}
