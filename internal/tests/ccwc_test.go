package tests

import (
	"bufio"
	"os"
	"testing"

	"github.com/wipdev-tech/ccwc/internal/counters"
)

func TestCountBytes(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal("err opening test file:", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	got := counters.Bytes(r)
	exp := 342190

	if exp != got {
		t.Fatalf("byte count failed: expected %v, got %v", exp, got)
	}
}
func TestCountLines(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal("err opening test file:", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	got := counters.Lines(r)
	exp := 7145

	if exp != got {
		t.Fatalf("line count failed: expected %v, got %v", exp, got)
	}
}

func TestCountWords(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal("err opening test file:", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	got := counters.Words(r)
	exp := 58164

	if exp != got {
		t.Fatalf("line count failed: expected %v, got %v", exp, got)
	}
}

func TestCountRunes(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal("err opening test file:", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	got := counters.Runes(r)
	exp := 339292

	if exp != got {
		t.Fatalf("line count failed: expected %v, got %v", exp, got)
	}
}

func TestFormatOutput(t *testing.T) {
	got := counters.FormatOutput("test.txt", 342190)
	exp := "  342190 test.txt"

	if exp != got {
		t.Fatalf("expected '%v', got '%v'", exp, got)
	}

	got = counters.FormatOutput("test.txt", 7145)
	exp = "    7145 test.txt"

	if exp != got {
		t.Fatalf("expected '%v', got '%v'", exp, got)
	}

	got = counters.FormatOutput("test.txt", 7145, 58164, 342190)
	exp = "    7145   58164  342190 test.txt"

	if exp != got {
		t.Fatalf("expected '%v', got '%v'", exp, got)
	}
}
