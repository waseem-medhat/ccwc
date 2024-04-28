package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wipdev-tech/ccwc/internal/counters"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	var fileName string
	var countBytes, countLines, countWords, countRunes bool
	if len(os.Args) == 2 {
		fileName = os.Args[1]
		countLines = true
		countWords = true
		countBytes = true
	}

	fileName = os.Args[len(os.Args)-1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, a := range os.Args {
		switch a {
		case "-c":
			countBytes = true
		case "-l":
			countLines = true
		case "-w":
			countWords = true
		case "-m":
			countRunes = true
		}
	}

	r := bufio.NewReader(f)

	results := []int{}

	if countLines {
		results = append(results, counters.Lines(r))
		f.Seek(0, 0)
	}

	if countWords {
		results = append(results, counters.Words(r))
		f.Seek(0, 0)
	}

	if countBytes {
		results = append(results, counters.Bytes(r))
		f.Seek(0, 0)
	}

	if countRunes {
		results = append(results, counters.Runes(r))
		f.Seek(0, 0)
	}

	fmt.Println(counters.FormatOutput(f.Name(), results...))
}
