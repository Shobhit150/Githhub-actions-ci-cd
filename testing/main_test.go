package main

import (
	"testing"
	"github.com/Shobhit150/Githhub-actions-ci-cd/function"
)

func TestDummy(t *testing.T) {
	t.Log("Starting TestDummy")
	got := function.Sum(3, 2)
	t.Logf("Sum(3, 2) returned %d", got)
	
	if got != 5 {
		t.Fatalf("Expected 5, got %d", got)
	}


	t.Log("TestDummy completed successfully")
}
