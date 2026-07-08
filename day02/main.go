// Day 2 — Types, zero values, constants, and iota.
//
// Run:   go run ./day02
//
// GOAL (≈20 min): model the status of a job in a distributed system as an
// enum-like type, using Go's `iota`. Go has NO `enum` keyword — the idiom is
// "a named integer type + a const block with iota". You'll build that, plus a
// String() method so the values print as names, not numbers.
//
// This is directly relevant to your work: job/task states, connection states,
// request phases — you model these constantly in distributed systems.
//
// ─────────────────────────────────────────────────────────────────────────
// TASKS — do them in order. Delete the panic lines as you implement each.
// ─────────────────────────────────────────────────────────────────────────

package main

import "fmt"

// TASK 1 — Define a named type for job status.
// In Go you make an enum by first declaring a new type based on int:
//
//     type Status int
//
// This gives you a distinct type (not just an alias) — the compiler will stop
// you from mixing a Status with a raw int by accident. Declare it below.

// TODO TASK 1: declare `type Status int` here

// TASK 2 — Declare the enum values with a const block + iota.
//
// `iota` is a counter that resets to 0 at the start of each const block and
// increments by 1 for each line. So this pattern:
//
//     const (
//         StatusPending Status = iota // 0
//         StatusRunning               // 1  (iota keeps going, type is inferred)
//         StatusSucceeded             // 2
//         StatusFailed                // 3
//     )
//
// IMPORTANT design point: the FIRST value (iota == 0) is the ZERO VALUE of the
// type. If you declare `var s Status`, it's StatusPending. That's often what
// you want — a freshly-created job defaults to "pending". Think about whether
// your zero value is a sensible default; here it is.
//
// Declare the four statuses below (Pending, Running, Succeeded, Failed).

// TODO TASK 2: const block with iota here

// TASK 3 — Give Status a String() method so it prints as a name.
//
// Right now, fmt would print a Status as its number (0,1,2,3). Implement the
// `Stringer` interface (we cover interfaces properly on Day 8, but you'll use
// one today): any type with a method `String() string` gets pretty-printed by
// fmt automatically.
//
// Signature (note the RECEIVER `(s Status)` before the name — this is how Go
// attaches a method to a type):
//
//     func (s Status) String() string {
//         switch s {
//         case StatusPending:
//             return "PENDING"
//         ... etc ...
//         default:
//             return "UNKNOWN"
//         }
//     }
//
// Implement it below.

// TODO TASK 3: func (s Status) String() string { ... }

func main() {
	// TASK 4 — Prove it works.
	//
	// a) Declare a job status WITHOUT initializing it, to see the zero value:
	//        var s Status
	//        fmt.Println("default status:", s)   // should print PENDING
	//
	// b) Print all four explicitly, e.g.:
	//        fmt.Println(StatusPending, StatusRunning, StatusSucceeded, StatusFailed)
	//
	// c) Print the underlying number of one, to see the int behind the name:
	//        fmt.Printf("StatusFailed as int = %d\n", StatusFailed)
	//        (%d forces the integer view even though String() exists)
	//
	// Expected output:
	//   default status: PENDING
	//   PENDING RUNNING SUCCEEDED FAILED
	//   StatusFailed as int = 3

	fmt.Println("TODO: implement TASK 4")
}
