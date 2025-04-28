package models

type Results struct {
	Words       int `json:"words"`
	Digits      int `json:"digits"`
	SpecialChar int `json:"special_char"`
	Lines       int `json:"lines"`
	Spaces      int `json:"spaces"`
	Sentences   int `json:"sentences"`
	Punctuation int `json:"punctuation"`
	Consonants  int `json:"consonants"`
	Vowels      int `json:"vowels"`
}
