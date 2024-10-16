package models

import "time"

type Word struct {
	ID      string `json:"id"`
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	Retries int    `json:"retries"`
	// TODO: Add field for last time the word was edited including the retry count
	CreatedAt      time.Time `json:"created_at"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}
