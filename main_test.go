package main

import (
	"os"
	"testing"
)

func TestGetPortSuccess(t *testing.T) {
	expected := ":8080"
	os.Setenv("PORT", "8080")
	result := getPort()
	if expected != result {
		t.Fatalf("expected: %s, actual: %s", expected, result)
	}
}
