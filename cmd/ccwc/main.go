package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wipdev-tech/ccwc/internal/counters"
)

func main() {
	f, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	switch os.Args[1] {
	case "-c":
		nBytes := counters.Bytes(r)
		fmt.Println(counters.FormatOutput(f.Name(), nBytes))
	case "-l":
		nLines := counters.Lines(r)
		fmt.Println(counters.FormatOutput(f.Name(), nLines))
	default:
		log.Fatal("malformed input")
	}
}
