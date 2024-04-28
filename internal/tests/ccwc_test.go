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

	r := bufio.NewReader(f)
	got := counters.Bytes(r)
	exp := 342190

	if exp != got {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}

func TestFormatOutput(t *testing.T) {
	got := counters.FormatOutput("test.txt", 342190)
	exp := "  342190 test.txt"

	if exp != got {
		t.Fatalf("expected '%v', got '%v'", exp, got)
	}
}
