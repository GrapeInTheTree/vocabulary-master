package models

type Word struct {
	ID   string `json:"id"`
	Word string `json:"word"`
	Meaning string `json:"meaning"`
	Retries int    `json:"retries"`
}