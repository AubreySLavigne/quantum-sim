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
	// Flip the 0x01 bit
	qs.value ^= 1
}

// Measure returns the value of this Quantum State
//
// If the state is determined, return its value.
//
// If it is undetermined, randomly return 0 or 1. This value is now
// determined, so Measuring the "undetermined" state will return the same
// value as the first
func (qs *QubitState) Measure() int {
	if qs.undetermined == false {
		return qs.value
	}

	// TODO: Implement the randomness for this
	return 0
}
