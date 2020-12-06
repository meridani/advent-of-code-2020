package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/meridani/advent-of-code-2020/pkg/days"
	"github.com/meridani/advent-of-code-2020/pkg/days/day1"
)

var listOfDays = make(map[int]string)
var currentDay = 1

func init() {
	for i := 1; i <= currentDay; i++ {
		listOfDays[i] = "./assets/" + string(i) + ".txt"
	}
}

func readInput(fileName string) (result []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Cannot open: ", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	for day, filename := range listOfDays {
		fmt.Println("Reading Day ", day, " input...")
		input := readInput(filename)
		today := &day1.Computer{}
	}
}
