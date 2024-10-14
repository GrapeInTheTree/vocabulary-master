package service

import (
	"github.com/google/uuid"
	vocaModels "github.com/grapeinthetree/vocabulary-master/internal/models"
	vocaRepository "github.com/grapeinthetree/vocabulary-master/internal/repository"
)

func StoreWord(word, meaning string) error {
	newWord := vocaModels.Word{
		ID:      uuid.New().String(),
		Word:    word,
		Meaning: meaning,
		Retries: 0,
	}
	return vocaRepository.InsertWord(newWord)
}

func RetrieveAllWords() ([]vocaModels.Word, error) {
	return vocaRepository.GetAllWords()
}

func RetrieveWordByName(word string) (vocaModels.Word, error) {
	return vocaRepository.GetWordByName(word)
}

func GetWordsForStudy(minRetryCount int) ([]vocaModels.Word, error) {
	return vocaRepository.GetWordsWithMinRetries(minRetryCount)
}

func GetWordsForExport(minRetryCount int) ([]vocaModels.Word, error) {
	return vocaRepository.GetWordsWithMinRetries(minRetryCount)
}

func UpdateWord(word, meaning string) error {
	return vocaRepository.UpdateWord(word, meaning)
}