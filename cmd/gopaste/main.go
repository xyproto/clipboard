package main

import (
	"fmt"

	"github.com/xyproto/clip"
)

func main() {
	text, err := clip.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Print(text)
}
