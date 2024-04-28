package main

import (
	"bufio"
	"fmt"
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
	nBytes := counters.Bytes(r)
	fmt.Println(counters.FormatOutput(f.Name(), nBytes, 3, 4))
}
