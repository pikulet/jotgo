# jotgo
jotto in golang! jotto is the game 'mastermind' applied to words (thank you enchant)

### Setup for jotgo game

If you want to develop jotgo applications using my engine, do

`sudo apt install libenchant-dev`

`go get -u github.com/pikulet/jotgo`

In your file, do

`import "github.com/pikulet/jotgo"`

Example usage

```
wordOne = "frisk"
wordTwo = "plane"
wordLength = 5

jotgo = game.CreateGame(wordLength)
jotgo.SetWord(game.PlayerOne, wordOne)
jotgo.SetWord(game.PlayerTwo, wordTwo)
```
