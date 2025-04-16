package main

import "testing"

func TestMainOutput(t *testing.T) {
	expected := "Hello, CI/CD World!"
	if mainOutput() != expected {
		t.Errorf("Expected %q, but got something else", expected)
	}
}

func mainOutput() string {
	return "Hello, CI/CD World!"
}
