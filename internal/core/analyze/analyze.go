package analyze

import "github.com/Ahmeds-Library/Go-Jwt/internal/models"

func Analyze(a string) models.Results {
	results := models.Results{}

	for i := 0; i < len(a); i++ {
		if a[i] == ' ' {
			results.Words++
			results.Spaces++
		} else if a[i] == '\n' {
			results.Words++
			results.Lines++
			results.Sentences++
		} else if a[i] == '\t' {
			results.Words++
		} else if a[i] == '.' || a[i] == ',' || a[i] == '!' || a[i] == '?' || a[i] == ';' || a[i] == ':' || a[i] == ')' || a[i] == '(' || a[i] == '-' || a[i] == '_' || a[i] == '"' || a[i] == '\'' {
			results.Punctuation++
		} else if a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U' {
			results.Vowels++
		} else if a[i] >= '0' && a[i] <= '9' {
			results.Digits++
		} else if a[i] == '@' || a[i] == '#' || a[i] == '$' || a[i] == '%' || a[i] == '^' || a[i] == '&' || a[i] == '*' || a[i] == '+' || a[i] == '=' || a[i] == '{' || a[i] == '}' || a[i] == '[' || a[i] == ']' {
			results.SpecialChar++
		} else if (a[i] >= 'a' && a[i] <= 'z' || a[i] >= 'A' && a[i] <= 'Z') && !(a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U') {
			results.Consonants++
		}
	}

	// Adjust for the last line and paragraph
	results.Lines++
	results.Words++

	return results
}
