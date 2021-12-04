package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var path string = "inputs/day-3.txt"
	//var path string = "inputs/day-3-test.txt"
	fmt.Println("Final part1 output: ", part1(path))
	//fmt.Println("Final part2 output: ", part2(path))
}

func part1(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	max_bit := 0
	var current string
	var bitcount [12]int
	var current_int int
	gamma, epsilon := "", ""

	for scanner.Scan() {
		max_bit++
		current = scanner.Text()

		for index, character := range current {
			current_int, _ = strconv.Atoi(string(character))
			bitcount[index] += current_int
		}
	}

	for _, count := range bitcount {
		if count > (max_bit - count) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return g * e
}
