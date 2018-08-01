package main

// QubitState represents the possible states that a Qubit may be.
// They may be 0, 1, or some combination of both.
type QubitState struct {
	value        int
	undetermined bool
}

// Represents that the value of the state is currently unkown
func undeterminedState() QubitState {
	qs := QubitState{}
	qs.undetermined = true
	return qs
}

// Returns a Qubitstate defaulting to 0.
// This will be the known bottom value of a newly created qubit
func defaultValue() QubitState {
	qs := QubitState{}
	qs.value = 0
	return qs
}

// Flip the qubit state
// 1 becomes 0 and vice-versa
func (qs *QubitState) Flip() {
	if qs.value == 0 {
		qs.value = 1
	} else {
		qs.value = 0
	}
}

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
