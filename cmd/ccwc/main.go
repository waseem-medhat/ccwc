package main

import (
	"bufio"
	"log"
	"os"

	"github.com/wipdev-tech/ccwc/internal/counters"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)
	counters.Bytes(r)
}
