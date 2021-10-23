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

func (raw *RawData) GetRawData() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		raw.RawInputList = append(raw.RawInputList, scanner.Text())
	}
}

type RawInput interface {
	GetRawData()
}
