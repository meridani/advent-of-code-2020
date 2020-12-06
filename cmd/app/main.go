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
		listOfDays[i] = fmt.Sprintf("./assets/inputs/%x.txt", i)
	}
}

func readInput(fileName string) (result days.Input) {
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
		result := today.Part1(input)
		fmt.Println(result)
		result = today.Part2(input)
		fmt.Println(result)
	}
}
