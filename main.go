package main

import (
	"bufio"
	"encoding/csv"
	"errors"
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
			correctAnswer, _ := toInt(record[5])

			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("insira um número inteiro")
	}
	return i, nil
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s\033[33m\n", index+1, question.Text)
		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}
		fmt.Println("Digite a sua resposta: ")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-2])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			fmt.Printf("\033[32m %s\033[0m\n", "Correcto!")
			g.Points += 10
		} else {
			fmt.Printf("\033[31m %s\033[0m\n", "Errado!")
		}
	}
}

func main() {
	game1 := &GameState{}

	go game1.ProcessCSV()

	game1.Init()

	game1.Run()

	fmt.Printf("Você ganhou %d pontos!\n", game1.Points)
}
