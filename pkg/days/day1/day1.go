package day1

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/meridani/advent-of-code-2020/pkg/days"
)

// Computer for Advent of Code
type Computer struct {
}

// Part1 of day 1
func (d *Computer) Part1(input days.Input) (res days.Result) {
	var numbers []int
	for _, line := range input {
		current, _ := strconv.Atoi(line)
		numbers = append(numbers, current)
	}

	sort.Ints(numbers)

	found := false
	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				found = true
				res = days.Result(fmt.Sprint(numbers[i] * numbers[j]))
				break
			}
		}
		if found {
			break
		}
	}
	return
}

// Part2 of day 1
func (d *Computer) Part2(input days.Input) (res days.Result) {
	var numbers []int
	for _, line := range input {
		current, _ := strconv.Atoi(line)
		numbers = append(numbers, current)
	}

	sort.Ints(numbers)

	found := false
	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					res = days.Result(fmt.Sprint(numbers[i] * numbers[j] * numbers[k]))
					found = true
					break
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	return res
}
