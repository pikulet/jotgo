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

// TargetWord to guess, contains the word and a set of its letters
type TWord struct {
	word		string
	wordSet		map[rune]bool
}

func getEmptyTWord() *TWord {
	return &TWord {
		word	:	"",
		wordSet :	make(map[rune]bool),
	}
}

func (tWord *TWord) reset() {
	tWord.word 		= ""
	tWord.wordSet	= make(map[rune]bool)
}

type Game struct {
	wordLength          int
	tWordOne			*TWord
	tWordTwo			*TWord
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
		tWordOne:           getEmptyTWord(),
		tWordTwo:           getEmptyTWord(),
		turn:               PlayerOne,
		dictionary:         d,
	}

	return game
}

func (g *Game) getTWord(player Player) *TWord { 
	if player {
		return g.tWordOne
	} else {
		return g.tWordTwo
	}
}

func (g *Game) verifyValid(word string) *TWord {
	if len(word) != g.wordLength {
		panic(fmt.Sprintf("InputError: Word is not of length %d", g.wordLength))
	}

	word = strings.ToLower(word)
	
	if g.dictionary.Check(word) {
		panic("InputError: Not a valid dictionary word")
	} 

	isUnique, setOfLetters := isWordContainsUniqueLetters(word)

	if !isUnique {
		panic("InputError: Word does not contain all unique letters")
	}

	return &TWord {
		word:		word,
		wordSet:	setOfLetters,
	}
}

func (g *Game) AddWord(player Player, word string) {
	if len(g.getTWord(player).word) != 0 {
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
	tWord := g.getTWord(!player)

	if guess == tWord.word {
		g.Reset()
		return true, g.wordLength
	} else {
		return false, getNumCommonLetters(guess, tWord.wordSet)
	} 
}

func (g *Game) Reset() {
	g.tWordOne.reset()
	g.tWordTwo.reset()
	g.turn = PlayerOne
}
