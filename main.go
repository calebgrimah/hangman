package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{
	"Zombie", "Expensive", "Kitten", "Dictionary", "Motivational", "Arrangement", "Pleiad", "Allegory", "Nation",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	fmt.Printf("target word : %s", targetWord)
	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only...")
			continue
		}
		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}
}

func getRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	targetWord := dictionary[rand.Intn(len(dictionary))]
	return targetWord
}

func getWordGuesssingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, character := range targetWord {
		if character == ' ' {
			result += ""
		} else if guessedLetters[character] == true {
			result += fmt.Sprintf("%c", character)
		} else {
			result += "_"
		}
		fmt.Println(" ")
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("/states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func readInput() string {
	fmt.Printf("> ")
	data, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(data)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuesssingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, character := range targetWord {
		if !guessedLetters[unicode.ToLower(character)] {
			return false
		}
	}
	return true
}
func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true
	return guessedLetters
}
