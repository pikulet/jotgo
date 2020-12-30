package game

import (
	"fmt"
	"strings"
	"github.com/hermanschaaf/enchant"
)

type Player int

const (
	PlayerOne Player = iota
	PlayerTwo
)

const (
	minWordLength = 4
	maxWordLength = 7
)

type TargetWord struct {
	word		string
	wordSet		map[rune]bool
}

type Game struct {
	wordLength          int
	tWordOne			TargetWord
	tWordTwo			TargetWord
	turn                Player
	winner              Player
	dictionary          *enchant.Enchant
}

func isWordContainsUniqueLetters(word string) (bool, map[rune]bool) {
	m := make(map[rune]bool)
	for _, char := range word {
		_, isPresent := m[char]
		if isPresent {
			return false, make(map[rune]bool)
		}

		m[char] = true
	}

	return true, m
}

func verifyValid(d *enchant.Enchant, word string, length int) TargetWord {
	if len(word) != length {
		panic(fmt.Sprintf("InputError: Word is not of length %d", length))
	}

	word = strings.ToLower(word)

	if d.Check(word) {
		panic("InputError: Not a valid dictionary word")
	} 

	isUnique, setOfLetters := isWordContainsUniqueLetters(word)

	if !isUnique {
		panic("InputError: Word does not contain all unique letters")
	}

	tWord := TargetWord {
		word:		word,
		wordSet:	setOfLetters,
	}

	return tWord
}

func CreateNewGame(wordOne string, wordTwo string, wordLength int) *Game {

	if wordLength < minWordLength {
		panic(fmt.Sprintf("EngineError: Word length should be at least %d", minWordLength)) 
	} else if wordLength > maxWordLength {
		panic(fmt.Sprintf("EngineError: Word length should be at most %d", maxWordLength))
	}

	d, _ := enchant.NewEnchant()
	d.LoadDict("en_US")

	tWordOne := verifyValid(d, wordOne, wordLength)
	tWordTwo := verifyValid(d, wordTwo, wordLength)

	game := &Game {
		wordLength:         wordLength,
		tWordOne:           tWordOne,
		tWordTwo:           tWordTwo,
		turn:               PlayerOne,
		dictionary:         d,
	}

	return game
}

func getTargetWord(g *Game, player Player) TargetWord {
	if player == PlayerOne {
		return g.tWordTwo
	} else {
		return g.tWordOne
	}
}

func getNumCommonLetters(word string, wordSet map[rune]bool) int {
	count := 0
	for _, char := range word {
		_, isPresent := wordSet[char]
		if isPresent {
			count += 1
		}
	} 

	return count
}

func MakeGuess(g *Game, player Player, guess string) (bool, int) { 

	if player != g.turn {
		panic("EngineError: Wrong player's turn")
	}

	guess = verifyValid(g.dictionary, guess, g.wordLength).word
	tWord := getTargetWord(g, player)

	if guess == tWord.word {
		return true, g.wordLength
	} else {
		return false, getNumCommonLetters(guess, tWord.wordSet)
	} 
}
