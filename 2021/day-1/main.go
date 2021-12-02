package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs/day-1.txt")
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
			prev = current
			i++
		} else {
			if current > prev {
				// fmt.Println("Prev, current: ", prev, ", ", current)
				increase_count++
				prev = current
			} else {
				prev = current
			}
		}
		fmt.Println(current)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Final increase count: ", increase_count)
}
