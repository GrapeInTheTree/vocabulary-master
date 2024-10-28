package service

import (
	"time"

	"github.com/google/uuid"
	vocaModels "github.com/grapeinthetree/vocabulary-master/internal/models"
	vocaRepository "github.com/grapeinthetree/vocabulary-master/internal/repository"
)

func StoreWord(word, meaning string) error {
	newWord := vocaModels.Word{
		ID:             uuid.New().String(),
		Word:           word,
		Meaning:        meaning,
		Retries:        0,
		CreatedAt:      time.Now(),
		LastModifiedAt: time.Now(),
	}
	return vocaRepository.InsertWord(newWord)
}

func RetrieveAllWords() ([]vocaModels.Word, error) {
	return vocaRepository.GetAllWords()
}

func RetrieveWordByName(word string) (vocaModels.Word, error) {
	return vocaRepository.GetWordByName(word)
}

func RetrieveWordsForToday() ([]vocaModels.Word, error) {
	return vocaRepository.GetWordsForToday()
}

func GetWordsForStudy(minRetryCount int) ([]vocaModels.Word, error) {
	return vocaRepository.GetWordsWithMinRetries(minRetryCount)
}

func GetWordsForExport(filename string, mode string, retryCount int) error {
	return vocaRepository.GetExportWordsToCSV(filename, mode, retryCount)
}

func UpdateWord(originalWord, newWord, meaning string) error {
	return vocaRepository.GetUpdateWord(originalWord, newWord, meaning)
}

func IncrementRetryCount(id string) error {
	return vocaRepository.GetIncrementRetryCount(id)
}
