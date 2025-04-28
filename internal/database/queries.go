package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Ahmeds-Library/Go-Jwt/internal/models"
)

const (
	SaveResultQuery = `INSERT INTO Results 
			(words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels, user_id, username)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
)

func GetUserCredentials(username string) (string, string, error) {
	var dbPassword, dbID string
	err := Db.QueryRow("SELECT password, id FROM users WHERE username = $1", username).Scan(&dbPassword, &dbID)
	if err == sql.ErrNoRows {
		return "", "", errors.New("user not found")
	}
	if err != nil {
		return "", "", err
	}
	return dbPassword, dbID, nil
}

func FetchResults(username, userID string, limit, offset int) ([]models.Results, error) {
	rows, err := Db.Query(`
        SELECT words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels
        FROM results
        WHERE username = $1 AND user_id = $2
        LIMIT $3 OFFSET $4`, username, userID, limit, offset)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func CountResults(username, userID string) (int, error) {
	var total int
	err := Db.QueryRow(`
        SELECT COUNT(*)
        FROM results
        WHERE username = $1 AND user_id = $2`, username, userID).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func CreateUser(username, password string) error {
	_, err := Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}
