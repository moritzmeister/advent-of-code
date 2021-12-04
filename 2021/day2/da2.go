package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var path string = "inputs/day-2.txt"
	//var path string = "inputs/day-2-test.txt"
	fmt.Println("Final part1 output: ", part1(path))
	fmt.Println("Final part2 output: ", part2(path))
}

func part2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontal, depth, aim, command_m := 0, 0, 0, 0
	var command string
	var current_split []string

	for scanner.Scan() {
		current_split = strings.Split(scanner.Text(), " ")
		//fmt.Println(current_split)
		command = current_split[0]
		command_m, _ = strconv.Atoi(current_split[1])

		if command == "forward" {
			horizontal += command_m
			depth += aim * command_m
		} else if command == "up" {
			aim -= command_m
		} else if command == "down" {
			aim += command_m
		} else {
			fmt.Println("error wrong command: ", command)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return horizontal * depth
}

func part1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontal, depth, command_m := 0, 0, 0
	var command string
	var current_split []string

	for scanner.Scan() {
		current_split = strings.Split(scanner.Text(), " ")
		//fmt.Println(current_split)
		command = current_split[0]
		command_m, _ = strconv.Atoi(current_split[1])

		if command == "forward" {
			horizontal += command_m
		} else if command == "up" {
			depth -= command_m
		} else if command == "down" {
			depth += command_m
		} else {
			fmt.Println("error wrong command: ", command)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return horizontal * depth
}
