package ditos

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Message struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
	Username     string `json:"username"`
	IconUrl      string `json:"icon_url"`
}

var ditos []string

func GenerateRandomDito() (*Message, error) {

	rand.Seed(time.Now().Unix())
	message := ditos[rand.Intn(len(ditos))]

	fmt.Println(message)

	var bah = &Message{
		ResponseType: "in_channel",
		Text:         message,
		Username:     "Gaucho Macho",
		IconUrl:      "http://www.eev.com.br/sipat/imagens/imgFotos65446.jpg",
	}

	return bah, nil
}

func LoadDitos() error {
	var err error
	ditos, err = readLines("ditos/ditos.txt")

	return err
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
