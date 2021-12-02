package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var path string = "../inputs/day-1.txt"
	//var tpath string = "../inputs/day-1-test.txt"
	fmt.Println("Final part1 increase count: ", part1(path))
	fmt.Println("Final part2 increase count: ", part2(path))
}

func sum(array [3]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func part2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	count := 0
	increase_count := 0
	var current, prev_sum, current_sum int
	var window [3]int

	for scanner.Scan() {
		current, _ = strconv.Atoi(scanner.Text())
		window[i] = current
		if i == 2 {
			i = 0
		} else {
			i++
		}
		current_sum = sum(window)
		if count > 2 {
			if current_sum > prev_sum {
				increase_count++
			}
		}
		count++
		//fmt.Println(current, window, prev_sum, current_sum, increase_count)
		prev_sum = current_sum
	}
	return increase_count
}

func part1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 1
	increase_count := 0
	var prev, current int

	for scanner.Scan() {
		current, _ = strconv.Atoi(scanner.Text())
		if i == 1 {
			i++
		} else {
			if current > prev {
				// fmt.Println("Prev, current: ", prev, ", ", current)
				increase_count++
			}
		}
		prev = current
		// fmt.Println(current)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return increase_count
}
