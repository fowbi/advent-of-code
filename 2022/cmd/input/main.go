package main

import (
	"aoc_2022/aoc"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func parseFlags() (day int, cookie string) {
	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "The AOC challenge of the day")
	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "Your AOC session cookie")
	flag.Parse()

	if day > 25 || day < 1 {
		log.Fatalf("Choose between 1 and 25: got %d", day)
	}

	if cookie == "" {
		log.Fatal("Cookie is missing")
	}

	return day, cookie
}

func main() {
	day, cookie := parseFlags()

	fmt.Printf("Fetching AOC input challenge for day %d ...", day)
	aocClient := aoc.AocApi()
	_, err := aocClient.SaveInput(day, cookie)

	if err != nil {
		log.Fatalf("\nCould not fetch input for day %d: %v", day, err)
	}

	fmt.Print(" [DONE]\n")
}
