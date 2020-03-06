package main

import (
	"fmt"

	"github.com/xyproto/clip"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(text)
}
