// Day 1 — Toolchain, modules, main, fmt.
//
// Run this file with:   go run ./day01
//
// GOAL (≈20 min): build a tiny temperature converter CLI.
// You already know this logic cold in Java/Python — that's the point.
// Today is purely about Go's shape: package, import, func main, fmt, and how
// values/printing work. Don't worry about anything "advanced."
//
// ─────────────────────────────────────────────────────────────────────────
// TASKS — do them in order. Delete the panic lines as you implement each.
// ─────────────────────────────────────────────────────────────────────────

package main

import "fmt"

// TASK 1 — Implement this function.
// It converts Celsius to Fahrenheit: F = C*9/5 + 32
// Note the Go signature style: parameter name THEN type, return type at the end.
func celsiusToFahrenheit(c float64) float64 {
	panic("TASK 1: not implemented yet")
}

// TASK 2 — Implement the reverse: C = (F - 32) * 5/9
func fahrenheitToCelsius(f float64) float64 {
	panic("TASK 2: not implemented yet")
}

func main() {
	// TASK 3 — Print a small conversion table for 0, 25, 37, 100 °C.
	//
	// Hints (Go specifics you're learning today):
	//   - A slice literal:      temps := []float64{0, 25, 37, 100}
	//   - Range loop:           for _, c := range temps { ... }
	//        ('_' discards the index — Go forbids unused variables, so you
	//         can't just write `for i, c := range` and ignore i.)
	//   - Formatted print:      fmt.Printf("%.1f°C = %.1f°F\n", c, f)
	//        (%.1f = float with 1 decimal; \n = newline; Printf does NOT add one.)
	//
	// Expected output:
	//   0.0°C = 32.0°F
	//   25.0°C = 77.0°F
	//   37.0°C = 98.6°F
	//   100.0°C = 212.0°F

	fmt.Println("TODO: print the conversion table")

	// TASK 4 (stretch) — pick one Fahrenheit value (say 98.6) and print it
	// converted back to Celsius, to sanity-check fahrenheitToCelsius.
}
