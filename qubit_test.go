package main

import "testing"

func assert(t *testing.T, q Qubit, expectedVal int, actions string) {
	res := q.Measure()
	if expectedVal != res {
		t.Errorf("'|0> %s [measure]' expects %d, got %d", actions, expectedVal, res)
	}
}

func TestInit(t *testing.T) {
	q := NewQubit()
	assert(t, q, 0, "")
}

func TestDoubleFlip(t *testing.T) {
	q := NewQubit()
	q.H().H()
	assert(t, q, 0, "H H")
}

func TestXInvert(t *testing.T) {
	q := NewQubit()
	q.X()
	assert(t, q, 1, "X")
}

func TestZInvert(t *testing.T) {
	q := NewQubit()
	q.H().Z().H()
	assert(t, q, 1, "H Z H")
}

func TestBottomAfterZ(t *testing.T) {
	q := NewQubit()
	q.Z()
	assert(t, q, 0, "Z")
}

func TestTopAfterX(t *testing.T) {
	q := NewQubit()
	q.H().X().H()
	assert(t, q, 0, "H X H")
}

func TestDoubleInvert(t *testing.T) {
	q := NewQubit()
	q.X().X()
	assert(t, q, 0, "X X")
}

func TestInvertBoth(t *testing.T) {
	q := NewQubit()
	q.Z().X()
	assert(t, q, 1, "Z X")
}
