package ditos

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Ditos struct {
	ditos []string
}

func New() (*Ditos, error) {
	ditos, err := readLines("ditos/ditos.txt")
	if err != nil {
		return nil, err
	}

	return &Ditos{
		ditos: ditos,
	}, nil
}

func (d *Ditos) Random() string {
	return d.ditos[rand.Intn(len(d.ditos))]
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
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
