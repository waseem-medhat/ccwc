package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	stdinStat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal("error checking stdin:", err)
	}

	fromStdin := stdinStat.Mode()&os.ModeNamedPipe == os.ModeNamedPipe

	var file *os.File
	var args []string
	var fileName string

	if fromStdin {
		file = os.Stdin
		if len(os.Args) > 1 {
			args = os.Args[1:]
		}
	} else if len(os.Args) < 2 {
		log.Fatal("no arguments")
	} else {
		fileName = os.Args[len(os.Args)-1]
		args = os.Args[1 : len(os.Args)-1]

		file, err = os.Open(fileName)
		if err != nil {
			log.Fatal("error opening file:", err)
		}
	}

	var printBytes, printRunes, printWords, printLines bool
	for _, a := range args {
		switch a {
		case "-c":
			printBytes = true
		case "-l":
			printLines = true
		case "-w":
			printWords = true
		case "-m":
			printRunes = true
		default:
			log.Fatal("malformed input")
		}
	}

	if len(args) == 0 {
		printBytes = true
		printLines = true
		printWords = true
	}

	nBytes, nRunes, nWords, nLines := count(file)
	var results []int

	if printLines {
		results = append(results, nLines)
	}

	if printWords {
		results = append(results, nWords)
	}

	if printBytes {
		results = append(results, nBytes)
	}

	if printRunes {
		results = append(results, nRunes)
	}

	fmt.Println(format(fileName, results...))
}

func count(file *os.File) (int, int, int, int) {
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	nBytes := len(bytes)
	nRunes := utf8.RuneCount(bytes)
	nWords := len(strings.Fields(string(bytes)))
	nLines := strings.Count(string(bytes), "\n")

	return nBytes, nRunes, nWords, nLines
}

func format(fileName string, ns ...int) string {
	lMax := len(fileName)
	for _, n := range ns {
		if l := len(fmt.Sprint(n)); l > lMax {
			lMax = l
		}
	}

	fmtStr := fmt.Sprintf("%%"+"%v"+"v", lMax)
	out := ""
	for _, n := range ns {
		if len(fmt.Sprint(n)) == lMax {
			out += " "
		}
		out += fmt.Sprintf(fmtStr, n)
	}

	out += " "
	out += fmt.Sprintf(fmtStr, fileName)
	return out
}
