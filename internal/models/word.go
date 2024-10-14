package models

type Word struct {
	ID      string `json:"id"`
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	Retries int    `json:"retries"`
	// TODO: Add field for last time the word was edited including the retry count
}
