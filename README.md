# :crayon: jotgo :crayon:

Jotto in golang! Jotto is the game 'mastermind' applied to words. For word-related dev, you would of course have to thank enchant.

## :firecracker: Usage

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
