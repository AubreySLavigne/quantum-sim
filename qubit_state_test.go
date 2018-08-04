package main

import "testing"

// Tests initialization of a default state
func TestDefaultInit(t *testing.T) {
	qs := defaultValue()
	if res := qs.Measure(); 0 != res {
		t.Errorf("'Expected %d, got %d", 0, res)
	}
}

// Tests flipping the value once
func TestFlipOnce(t *testing.T) {
	qs := defaultValue()
	qs.Flip()
	if res := qs.Measure(); 1 != res {
		t.Errorf("'Expected %d, got %d", 1, res)
	}
}

// Tests flipping the value twice
func TestFlipTwice(t *testing.T) {
	qs := defaultValue()
	qs.Flip()
	qs.Flip()
	if res := qs.Measure(); 0 != res {
		t.Errorf("'Expected %d, got %d", 0, res)
	}
}
