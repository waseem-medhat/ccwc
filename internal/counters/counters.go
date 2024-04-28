package counters

import (
	"bufio"
	"fmt"
)

func Bytes(r *bufio.Reader) int {
	nBytes := 0
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)
	for s.Scan() {
		nBytes += len(s.Text())
	}
	return nBytes
}

func FormatOutput(fileName string, ns ...int) string {
	lMax := len(fileName)
	for _, n := range ns {
		if l := len(fmt.Sprint(n)); l > lMax {
			lMax = l
		}
	}

	fmt.Println(lMax)

	fmtStr := fmt.Sprintf("%%"+"%v"+"v", lMax)
	out := ""
	for _, n := range ns {
		out += fmt.Sprintf(fmtStr, n)
		out += " "
	}

	out += fmt.Sprintf(fmtStr, fileName)
	return out
}
