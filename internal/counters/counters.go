package counters

import "bufio"

func Bytes(r *bufio.Reader) int {
	nBytes := 0
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)
	for s.Scan() {
		nBytes += len(s.Text())
	}
	return nBytes
}
