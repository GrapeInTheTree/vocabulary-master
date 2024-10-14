package repository

import (
	"database/sql"
	"fmt"
	vocaModels "github.com/grapeinthetree/vocabulary-master/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./vocabulary.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	err = createTable()
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}

func createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS words (
		id TEXT PRIMARY KEY,
		word TEXT NOT NULL,
		meaning TEXT NOT NULL,
		retries INTEGER DEFAULT 0
	);`

	_, err := DB.Exec(query)
	return err
}

func InsertWord(word vocaModels.Word) error {
	query := `INSERT INTO words (id, word, meaning, retries) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, word.ID, word.Word, word.Meaning, word.Retries)
	return err
}

func GetAllWords() ([]vocaModels.Word, error) {
	query := `SELECT id, word, meaning, retries FROM words`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []vocaModels.Word
	for rows.Next() {
		var w vocaModels.Word
		err := rows.Scan(&w.ID, &w.Word, &w.Meaning, &w.Retries)
		if err != nil {
			return nil, err
		}
		words = append(words, w)
	}
	return words, nil
}

func GetWordByName(word string) (vocaModels.Word, error) {
	query := `SELECT id, word, meaning, retries FROM words WHERE word = ?`
	row := DB.QueryRow(query, word)
	var w vocaModels.Word
	err := row.Scan(&w.ID, &w.Word, &w.Meaning, &w.Retries)
	return w, err
}

func GetWordsWithMinRetries(minRetries int) ([]vocaModels.Word, error) {
	query := `SELECT id, word, meaning, retries FROM words WHERE retries > ?`
	rows, err := DB.Query(query, minRetries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []vocaModels.Word
	for rows.Next() {
		var w vocaModels.Word
		err := rows.Scan(&w.ID, &w.Word, &w.Meaning, &w.Retries)
		if err != nil {
			return nil, err
		}
		words = append(words, w)
	}
	return words, nil
}

func UpdateWord(word, meaning string) error {
	query := `UPDATE words SET meaning = ? WHERE word = ?`
	_, err := DB.Exec(query, meaning, word)
	return err
}
