# competitiveprogtemp-go

Easy Pack for Competitive Programming Beginners

sample

```Go
package main

import (
	"fmt"
    "os"
	"github.com/Iovesophy/competitiveprogtemp-go/easy"
)

func main() {
    os.Stdin = "test"
	e := easy.RawInput{}
	e.GetRawData()
    fmt.Println(e.RawInputList[0])
}

```

License
Copyright (c) 2021 Kazuya yuda. This software is released under the MIT License, see LICENSE. https://opensource.org/licenses/mit-license.php