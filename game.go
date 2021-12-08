package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
)

type Game struct {
	Word             []rune
	guessingWord     []rune
	allowedMissCount int
	scanner          *bufio.Scanner
	Out              chan string
	In               chan rune
	Win              chan bool
}

func NewGame(wordList []string, allowedMissesCount int, scanner *bufio.Scanner) *Game {
	word := []rune(wordList[rand.Intn(len(wordList))])

	guessingWord := make([]rune, len(word))
	for i := range guessingWord {
		guessingWord[i] = '_'
	}
	return &Game{
		Word:             word,
		guessingWord:     guessingWord,
		allowedMissCount: allowedMissesCount,
		scanner:          scanner,
		Out:              make(chan string),
		In:               make(chan rune),
		Win:              make(chan bool),
	}
}

func (game *Game) Play() (win bool, err error) {
	missCount := 0

	for !game.wordComplete() && missCount <= game.allowedMissCount {
		game.printGuessWord()
		guess, valid := game.readPlayerInput()
		if !valid {
			fmt.Println("Please enter a valid guess")
			continue
		}

		validGuess := game.checkGuess(guess)

		if !validGuess {
			missCount++
		}

		if err := game.scanner.Err(); err != nil {
			panic(err)
			//return false, err
		}
	}

	game.printGuessWord()
	return missCount <= game.allowedMissCount, nil
}

func (game *Game) printGuessWord() {
	fmt.Println("\n" + string(game.guessingWord))
}

func (game *Game) readPlayerInput() (guess rune, valid bool) {
	fmt.Print("Guess: ")

	game.scanner.Scan()
	runes := []rune(game.scanner.Text())
	if len(runes) <= 0 {
		return ' ', false
	}
	return runes[0], true
}

func (game *Game) checkGuess(guess rune) (validGuess bool) {
	validGuess = false

	for i, char := range game.Word {
		if char == guess {
			game.guessingWord[i] = char
			validGuess = true
		}
	}

	return validGuess
}

func (game *Game) wordComplete() bool {
	for _, c := range game.guessingWord {
		if c == '_' {
			return false
		}
	}

	return true
}
