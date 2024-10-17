package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"time"

	vocaModels "github.com/grapeinthetree/vocabulary-master/internal/models"
	vocaRepository "github.com/grapeinthetree/vocabulary-master/internal/repository"
	vocaService "github.com/grapeinthetree/vocabulary-master/internal/service"
	"github.com/urfave/cli/v2"
)

func main() {
	err := vocaRepository.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	app := &cli.App{
		Name:  "vocabulary-cli",
		Usage: "A CLI vocabulary application",
		Commands: []*cli.Command{
			{
				Name:    "store",
				Aliases: []string{"s"},
				Usage:   "Store new words",
				Action:  storeWords,
			},
			{
				Name:    "retrieve",
				Aliases: []string{"r"},
				Usage:   "Retrieve stored words",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "Retrieve all words",
					},
				},
				Action: retrieveWords,
			},
			{
				Name:    "study",
				Aliases: []string{"st"},
				Usage:   "Study words (use --all or --only-retry)",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "only-retry",
						Aliases: []string{"r"},
						Usage:   "Only study words with more retries than this",
					},
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "Study all words",
					},
					&cli.BoolFlag{
						Name:    "today",
						Aliases: []string{"t"},
						Usage:   "Study words for today",
					},
				},
				Action: studyWords,
			},
			{
				Name:    "export",
				Aliases: []string{"e"},
				Usage:   "Export words",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "only-retry",
						Aliases: []string{"r"},
						Usage:   "Only export words with more retries than this",
					},
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "Export all words",
					},
				},
				Action: exportWords,
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update the database",
				Action:  updateWord,
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: add logic that if the argument is zero, then it will not store and throw an error
func storeWords(c *cli.Context) error {
	fmt.Println("Enter words and meanings to store (type 'exit' to finish):")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Word: ")
		if !scanner.Scan() {
			break
		}
		word := scanner.Text()
		if word == "exit" {
			break
		}
		fmt.Print("Meaning: ")
		if !scanner.Scan() {
			break
		}
		meaning := scanner.Text()
		err := vocaService.StoreWord(word, meaning)
		if err != nil {
			return err
		}
		fmt.Printf("Stored: %s - %s\n", word, meaning)
	}
	return scanner.Err()
}

func retrieveWords(c *cli.Context) error {
	if c.Bool("all") {
		words, err := vocaService.RetrieveAllWords()
		if err != nil {
			return err
		}
		for _, word := range words {
			fmt.Printf("%s - %s (Retries: %d)\n", word.Word, word.Meaning, word.Retries)
		}
	} else {
		word, err := vocaService.RetrieveWordByName(c.Args().First())
		if err != nil {
			return err
		}
		fmt.Printf("%s - %s (Retries: %d)\n", word.Word, word.Meaning, word.Retries)
	}
	return nil
}

// TODO: add logic to increase retry count if the word is not known
func studyWords(c *cli.Context) error {
	if !c.Bool("all") && !c.IsSet("only-retry") && !c.Bool("today") {
		return fmt.Errorf("please specify either --all or --only-retry or --today option")
	}
	var words []vocaModels.Word
	var err error

	if c.Bool("all") {
		words, err = vocaService.RetrieveAllWords()
	} else if c.Bool("today") {
		// TODO: implement logic to study words for today
		words, err = vocaService.RetrieveWordsForToday()
	} else {
		retryCount := c.Int("only-retry")
		words, err = vocaService.GetWordsForStudy(retryCount)
		fmt.Printf("Studying words with retry count > %d\n", retryCount)
	}

	if err != nil {
		return err
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return studyWordList(words)
}

func studyWordList(words []vocaModels.Word) error {
	for _, word := range words {
		fmt.Printf("Word: %s\n", word.Word)
		fmt.Print("Press Enter to see the meaning...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Printf("Meaning: %s\n", word.Meaning)
		fmt.Print("Press Enter to continue or 'w' to skip...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	return nil
}

func exportWords(c *cli.Context) error {
	if !c.Bool("all") && !c.IsSet("only-retry") {
		return fmt.Errorf("please specify either --all or --only-retry option")
	}

	if c.Bool("all") {
		return vocaService.GetWordsForExport(fmt.Sprintf("words-%s.csv", time.Now().Format("01-02")), "all", 0)
	}
	retryCount := c.Int("only-retry")
	return vocaService.GetWordsForExport(fmt.Sprintf("words-%s-retry-%d.csv", time.Now().Format("01-02"), retryCount), "retry", retryCount)
}

func updateWord(c *cli.Context) error {
	originalWord := c.Args().First()
	newWord := c.Args().Get(1)
	meaning := c.Args().Get(2)
	if originalWord == "" || meaning == "" {
		return fmt.Errorf("please provide a word and meaning to update")
	}
	return vocaService.UpdateWord(originalWord, newWord, meaning)
}
