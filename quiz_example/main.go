package main

import (
	"os"
	"fmt"
	"flag"
	"encoding/csv"
	"time"
)

func main() {
	filename := flag.String("file", "problems.csv", "File to read from")
	timer_sec := flag.Int("limit", 30, "Timer duration in seconds.")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		fmt.Printf("Error opening file with name %s\n", *filename)
		os.Exit(1)
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Printf("Error reading file with name %s\n", *filename)
	}

	timer := time.NewTimer(time.Duration(*timer_sec) * time.Second)
	correct_count := 0

quizloop:
	for i := 0; i < len(lines); i++ {
		fmt.Printf("What's %s? ", lines[i][0])

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time ran out.")
			break quizloop
		case result := <-answerCh:
			if result == lines[i][1] {
				correct_count++
			}
		}
	}

	fmt.Printf("Correct answer count: %d.\n", correct_count)
}
