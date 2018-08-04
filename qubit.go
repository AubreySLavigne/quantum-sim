package main

// Qubit represents a single quantum bit.
// These have a bottom and a top spin
type Qubit struct {
	top    QubitState
	bottom QubitState
}

// Measure reads the value of bottom
// If the qubit is undetermined, it chooses a 1 or 0 at random.
// TODO: Randomly choose a value if it is undetermined
func (q *Qubit) Measure() int {
	return q.bottom.value
}

// NewQubit returns a Qubit with default settings
func NewQubit() Qubit {
	q := Qubit{}
	q.top = undeterminedState()
	q.bottom = defaultValue()
	return q
}

// H gate swaps the bottom and top status of the qubit
func (q *Qubit) H() *Qubit {
	q.top, q.bottom = q.bottom, q.top
	return q
}

// X gate flips the bottom state
func (q *Qubit) X() *Qubit {
	q.bottom.Flip()
	return q
}

// Z gate flips the top state
func (q *Qubit) Z() *Qubit {
	q.top.Flip()
	return q
}
