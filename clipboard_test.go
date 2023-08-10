package clip_test

import (
	"testing"

	. "github.com/xyproto/clip"
)

func TestCopyAndPaste(t *testing.T) {
	expected := "日本語"

	err := WriteAll(expected, false)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAll(false)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestCopyAndPasteBytes(t *testing.T) {
	expected := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	err := WriteAllBytes(expected, false)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAllBytes(false)
	if err != nil {
		t.Fatal(err)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// For testing the test
	// expected[0] = 42

	for i := 0; i < min(len(actual), len(expected)); i++ {
		if actual[i] != expected[i] {
			t.Errorf("want %s, got %s", expected, actual)
			break
		}
	}
}

func TestMultiCopyAndPaste(t *testing.T) {
	expected1 := "French: éèêëàùœç"
	expected2 := "Weird UTF-8: 💩☃"

	err := WriteAll(expected1, false)
	if err != nil {
		t.Fatal(err)
	}

	actual1, err := ReadAll(false)
	if err != nil {
		t.Fatal(err)
	}
	if actual1 != expected1 {
		t.Errorf("want %s, got %s", expected1, actual1)
	}

	err = WriteAll(expected2, false)
	if err != nil {
		t.Fatal(err)
	}

	actual2, err := ReadAll(false)
	if err != nil {
		t.Fatal(err)
	}
	if actual2 != expected2 {
		t.Errorf("want %s, got %s", expected2, actual2)
	}
}

func BenchmarkReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAll(false)
	}
}

func BenchmarkWriteAll(b *testing.B) {
	text := "いろはにほへと"
	for i := 0; i < b.N; i++ {
		WriteAll(text, false)
	}
}

func BenchmarkReadAllBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAllBytes(false)
	}
}

var bs = []byte("いろはにほへと")

func BenchmarkWriteAllBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteAllBytes(bs, false)
	}
}
