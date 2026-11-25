package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("could not open file:", err)
		return
	}
	defer file.Close()

	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i) // turn 142 into the text "142"
		half := len(s) / 2

		fmt.Fprint(file, s[:half])        // write first part, e.g. "1"
		time.Sleep(50 * time.Millisecond) // crash can happen HERE
		fmt.Fprint(file, s[half:]+"\n")   // write the rest, e.g. "42"
	}
	fmt.Println("done")
}
