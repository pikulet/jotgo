package game

import (
	"fmt"
	"strings"
	"github.com/hermanschaaf/enchant"
)

type Player bool

const (
	PlayerOne Player = true
	PlayerTwo Player = false
)

const (
	minWordLength = 4
	maxWordLength = 7
)

type TargetWord struct {
	word		string
	wordSet		map[rune]bool
}

type Game struct {$
	wordLength          int
	tWordOne			TargetWord
	tWordTwo			TargetWord
	turn                Player
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


func CreateGame(wordLength int) *Game {

	if wordLength < minWordLength {
		panic(fmt.Sprintf("EngineError: Word length should be at least %d", minWordLength)) 
	} else if wordLength > maxWordLength {
		panic(fmt.Sprintf("EngineError: Word length should be at most %d", maxWordLength))
	}

	d, _ := enchant.NewEnchant()
	d.LoadDict("en_US")

	game := &Game {
		wordLength:         wordLength,
		tWordOne:           nil,
		tWordTwo:           nil,
		turn:               PlayerOne,
		dictionary:         d,
	}

	return game
}

func (g *Game) getWord(player Player) TargetWord { 
	if player {
		return g.tWordOne
	} else {
		return g.tWordTwo
	}
}

func (g *Game) verifyValid(word string) TargetWord {
	if len(word) != g.wordLength {
		panic(fmt.Sprintf("InputError: Word is not of length %d", length))
	}

	word = strings.ToLower(word)
	
	if g.dictionary.Check(word) {
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

func (g *Game) AddWord(player Player, word string) {
	isExists := g.getWord(player)
	if isExists {
		panic("InputError: Player has already set a word")
	}

	tWord := g.verifyValid(word)

	if player {
		g.tWordOne = tWord
	} else {
		g.tWordTwo = tWord
	}
}

func (g *Game) MakeGuess(player Player, guess string) (bool, int) { 

	if player != g.turn {
		panic("EngineError: Wrong player's turn")
	}

	// flip turn
	g.turn = !g.turn

	guess = g.verifyValid(guess).word
	tWord := g.getTargetWord(!player)

	if guess == tWord.word {
		g.Reset()
		return true, g.wordLength
	} else {
		return false, getNumCommonLetters(guess, tWord.wordSet)
	} 
}

func (g *Game) Reset() {
	g.tWordOne = nil
	g.tWordTwo = nil
	g.turn = PlayerOne
}
