package main

import (
	"fmt"
	"os"
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
		fmt.Fprintf(file, "%d\n", i)

		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("done")
}
