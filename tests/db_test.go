package tests

import (
	"testing"
)

func TestDB(t *testing.T) {
	_, err := "test", (*error)(nil)
	if err != nil {
		t.Fatal(err)
	}
}
