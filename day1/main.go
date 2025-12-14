package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadData() []string {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Yeblo!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	return data
}

func main() {
	data := loadData()

	currentPoint := 50
	dialSize := 100

	anwser := 0

	for i := 0; i < len(data); i++ {
		row := data[i]
		row = strings.TrimSpace(row)
		direction := row[:1]
		step, _ := strconv.Atoi(row[1:])
		if direction == "L" {
			step *= -1
		}
		fmt.Printf("dir %s size %d \n", direction, step)

		currentPoint = (currentPoint + step) % dialSize
		if currentPoint == 0 {
			anwser++
		}
	}

	fmt.Println(anwser)
}
