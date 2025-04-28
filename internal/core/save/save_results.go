package save_results

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Ahmeds-Library/Go-Jwt/internal/models"
)

func SaveResult(db *sql.DB, result models.Results, userID string, username string) error {
	query := `
			INSERT INTO Results 
			(words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels, user_id, username)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

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
