package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("BROKEN: cannot open data.txt")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expected := 1

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil || num != expected {
			fmt.Println("BROKEN at line", expected, "found:", line)
			os.Exit(1)
		}
		expected++
	}

	fmt.Println("OK,", expected-1, "lines")
}
