package easy

import (
	"bufio"
	"os"
)

type RawData struct {
	RawInputList []string
}

var (
	_ RawInput = (*RawData)(nil)
)

func (raw *RawData) GetRawData(n int) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		raw.RawInputList = append(raw.RawInputList, scanner.Text())
	}
}

type RawInput interface {
	GetRawData(int)
}
