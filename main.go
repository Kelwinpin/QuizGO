package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Seja Bem-vindo ao QuizGO!")
	fmt.Println("Qual o nome do jogador?")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	g.Name = name

	fmt.Printf("Vamos começar o jogo! %s\n", g.Name)
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quiz.csv")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for index, record := range records {
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	game1 := &GameState{}

	go game1.ProcessCSV()

	game1.Init()
}
