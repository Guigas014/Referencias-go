package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type GameState struct {
	words []string
	chosenWord string
	showWord string
}

//Body da request
type Request struct {
	Letter string `json:"letter"`
}

var gameState GameState

//Função que retorna a palavra escolhida.
func getWord(c echo.Context) error  {
	wrd := fmt.Sprintf("word: %s", gameState.showWord)
	return c.String(http.StatusOK, wrd)
}

//Função que processa a letra digitada e retorna a situação do jogo.
func insertWord(c echo.Context) error  {
	var req Request

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	processLetter(req.Letter)

	//Teste se todas as letras já foram descobertas e tira a palavra "word" que antecede a string de saída. Esse IF é o mesmo da função processLetter.
	if !strings.Contains(gameState.showWord, "_") {
		word := fmt.Sprintf(gameState.showWord)
		
		return c.String(http.StatusOK, word)
	}

	word := fmt.Sprintln("word: ", gameState.showWord)
	return c.String(http.StatusOK, word)
}

//Função que reinicia o jogo.
func newGame(c echo.Context) error {
	gameState.showWord = ""

	initGame()

	word := fmt.Sprintf("word: %s", gameState.showWord)
	return c.String(http.StatusOK, word)
}

func main() {
	initGame()

	e := echo.New()

	// CORS restricted
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/word", getWord)
	e.GET("/word/new", newGame)
	e.POST("/word", insertWord)

	e.Logger.Fatal(e.Start(":1323"))
}

//Função que inicia o jogo sorteando a palavra
func initGame() {
	gameState.words = []string {
		"peteca",
		"luz",
		"janela",
		"abacaxi",
		"suquinho",
	}

	// Gera indice aleatório
	reader := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := reader.Intn(len(gameState.words))

	gameState.chosenWord = gameState.words[index]

	for range gameState.chosenWord {
		gameState.showWord += "_"
	}
}

//Função que processa a letra digitada
func processLetter(letter string)  {
	letter = strings.ToLower(letter)

	if strings.Contains(gameState.chosenWord, letter) {
		for i, c := range gameState.chosenWord {
			if string(c) == letter {
				gameState.showWord = gameState.showWord[:i] + letter +  gameState.showWord[i+1:]
			}
		}
	}

	if !strings.Contains(gameState.showWord, "_") {
		txt := "Parabéns! Você acertou a palavra!!\n"
		txt += "A palavra era: " + gameState.chosenWord
		gameState.showWord = txt
	}
}
