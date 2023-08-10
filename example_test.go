package clip_test

import (
	"fmt"

	"github.com/xyproto/clip"
)

func Example() {
	clip.WriteAll("日本語", false)
	text, _ := clip.ReadAll(false)
	fmt.Println(text)

	// Output:
	// 日本語
}
