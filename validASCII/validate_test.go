package validASCII

import (
	"testing"
)

func TestAllASCII(t *testing.T) {
	input := "abc"
	want := "abc"
	result := Mark(input, Standard)

	if result != want {
		t.Errorf("Expected %s, got %s", want, result)
	}
}

func TestOneUnicode(t *testing.T) {
	input := "ab√"
	want := "ab�"
	result := Mark(input, Standard)

	if result != want {
		t.Errorf("Expected %s, got %s", want, result)
	}
}
