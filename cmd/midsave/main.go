package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	app := flag.String("app", "", "app to crash-test")
	check := flag.String("check", "", "program that checks the data")
	clean := flag.String("clean", "", "file to delete before each run")
	runs := flag.Int("runs", 10, "how many times to crash it")
	flag.Parse()

	failed := 0
	for i := 1; i <= *runs; i++ {
		os.Remove(*clean) // fresh start each run

		cmd := exec.Command(*app)
		cmd.Start() // 1. start the app

		wait := time.Duration(200+rand.Intn(2000)) * time.Millisecond
		time.Sleep(wait) // 2. wait a random time

		cmd.Process.Kill() // 3. kill it
		cmd.Wait()

		err := exec.Command(*check).Run() // 4. run the checker
		if err != nil {
			failed++
			fmt.Printf("run %d: BROKEN (killed after %v)\n", i, wait)
		} else {
			fmt.Printf("run %d: ok\n", i)
		}
	}
	fmt.Printf("\n%d of %d runs lost or broke data\n", failed, *runs)
}
