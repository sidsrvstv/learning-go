// Day 2 — Types, zero values, constants, and iota.
//
// Run:   go run ./day02
//
// Model a distributed-system job's status as an enum-like type using iota,
// with a String() method so it prints as a name instead of a number.

package main

import "fmt"

// A distinct named type based on int — the compiler won't let a raw int be
// used where a Status is expected.
type Status int

// iota resets to 0 at the top of the block and increments by 1 per line.
// The first value (0) is the type's ZERO VALUE, so `var s Status` == StatusPending.
const (
	StatusPending   Status = iota // 0
	StatusRunning                 // 1
	StatusSucceeded               // 2
	StatusFailed                  // 3
)

// Implementing String() (the Stringer interface) makes fmt print the name.
func (s Status) String() string {
	switch s {
	case StatusPending:
		return "PENDING"
	case StatusRunning:
		return "Running"
	case StatusSucceeded:
		return "Succeeded"
	case StatusFailed:
		return "Failed"
	default:
		return "UNKNOWN"
	}
}

func main() {
	// a) zero value — uninitialized Status defaults to StatusPending
	var s Status
	fmt.Println("default status:", s)

	// b) all four, printed by name via String()
	fmt.Println(StatusPending, StatusRunning, StatusSucceeded, StatusFailed)

	// c) the int behind the name — %d forces the integer view
	fmt.Printf("StatusFailed as int = %d\n", StatusFailed)
}
