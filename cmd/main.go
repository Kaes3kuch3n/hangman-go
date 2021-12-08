package main

import (
	"bufio"
	"fmt"
	"hangman"
	"math/rand"
	"os"
	"time"
)

var (
	scanner       = bufio.NewScanner(os.Stdin)
	allowedMisses = 10
)

func checkError(err error) {
	if err != nil {
		fmt.Println("File Read Error:", err)
		os.Exit(1)
	}
}

func readWordsFile(filePath string) []string {
	file, err := os.Open(filePath)
	checkError(err)
	defer func(file *os.File) {
		err := file.Close()
		checkError(err)
	}(file)

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	checkError(scanner.Err())
	return words
}

func main() {
	rand.Seed(time.Now().UnixNano())
	words := readWordsFile("assets/words.txt")
	game := hangman.NewGame(words, allowedMisses, scanner)
	win, err := game.Play()

	if err != nil {
		fmt.Println("Game Error:", err)
		os.Exit(2)
	}

	if win {
		fmt.Println("Yay, you won!")
		return
	}

	fmt.Println("You lost :(")
	fmt.Println("The word was", string(game.Word))
}
