// Package day01 drills Go's basic syntax: types, control flow, and functions.
//
// Every function below panics with "TODO". Replace each body with a real
// implementation until `go test ./day01/` passes. Work top to bottom — they
// get harder.
package day01

import "errors"

// ErrDivideByZero is a "sentinel error": a package-level error value that
// callers can compare against with errors.Is. You'll meet these properly on
// Day 3, but SafeDivide needs one now.
var ErrDivideByZero = errors.New("division by zero")

// ---------------------------------------------------------------------------
// 1. Control flow
// ---------------------------------------------------------------------------

// FizzBuzz returns "Fizz" if n is divisible by 3, "Buzz" if by 5, "FizzBuzz"
// if by both, and the decimal form of n otherwise.
//
// Hints: strconv.Itoa converts an int to a string. Note that Go has no ternary
// operator — use a plain if/else chain, or a `switch` with no condition, which
// is Go's idiomatic replacement for if/else-if ladders.
func FizzBuzz(n int) string {
	panic("TODO")
}

// Grade maps a 0–100 score to a letter: 90+ is "A", 80+ "B", 70+ "C", 60+ "D",
// below 60 "F". A score outside 0–100 returns "invalid".
//
// Hint: a conditionless `switch { case score >= 90: ... }` reads far better
// than if/else here. Go's switch does NOT fall through by default, so you do
// not need `break`.
func Grade(score int) string {
	panic("TODO")
}

// ---------------------------------------------------------------------------
// 2. Loops
// ---------------------------------------------------------------------------

// SumTo returns 1 + 2 + ... + n. For n < 1 it returns 0.
//
// Hint: Go has exactly one loop keyword, `for`. The C-style three-clause form
// is `for i := 1; i <= n; i++ {}`.
func SumTo(n int) int {
	panic("TODO")
}

// CollatzSteps returns how many steps it takes to reach 1 from n using the
// Collatz rule: if even, halve it; if odd, 3n+1. CollatzSteps(1) is 0.
// For n < 1 return -1.
//
// Hint: this is the `for condition {}` form — Go's while loop.
func CollatzSteps(n int) int {
	panic("TODO")
}

// ---------------------------------------------------------------------------
// 3. Strings, bytes, and runes  <-- the part that surprises people
// ---------------------------------------------------------------------------

// ReverseString returns s with its characters in reverse order.
//
// IMPORTANT: a Go string is a read-only slice of BYTES, not characters.
// Indexing s[i] gives you a byte. Ranging over a string with
// `for i, r := range s` gives you runes (Unicode code points) and the byte
// offset of each. Reversing byte-by-byte corrupts any non-ASCII text, and the
// tests check that — convert to []rune first.
func ReverseString(s string) string {
	panic("TODO")
}

// CountVowels returns the number of vowels (a, e, i, o, u — either case) in s.
//
// Hint: `for _, r := range s` yields each rune. strings.ContainsRune is handy.
func CountVowels(s string) int {
	panic("TODO")
}

// ---------------------------------------------------------------------------
// 4. Multiple return values
// ---------------------------------------------------------------------------

// MinMax returns the smallest and largest values in nums. For an empty slice
// it returns (0, 0, false); otherwise the third value is true.
//
// Returning a bool as the last value to signal "was there anything?" is a
// common Go pattern — the same shape as the comma-ok idiom on maps.
func MinMax(nums []int) (minimum, maximum int, ok bool) {
	panic("TODO")
}

// SafeDivide returns a/b, or 0 and ErrDivideByZero when b is 0.
//
// This is THE Go convention: errors are ordinary return values, not thrown.
// The caller must deal with the error explicitly.
func SafeDivide(a, b int) (int, error) {
	panic("TODO")
}

// ---------------------------------------------------------------------------
// 5. Variadic functions
// ---------------------------------------------------------------------------

// Average returns the mean of nums. With no arguments it returns 0.
//
// Hint: the signature `nums ...float64` makes nums a []float64 inside the
// function. Callers can also pass an existing slice with `Average(xs...)`.
func Average(nums ...float64) float64 {
	panic("TODO")
}

// ---------------------------------------------------------------------------
// 6. defer
// ---------------------------------------------------------------------------

// TraceOrder appends "start" to the returned slice, then uses THREE separate
// defer statements to append "one", "two", and "three" in that source order,
// then appends "end" before returning.
//
// The point of the exercise: deferred calls run LIFO (last deferred, first
// run), and they run after the return value is set — which is why this needs
// a NAMED return value (`out`) that the deferred closures can mutate.
//
// Work out the expected order yourself before reading the test.
func TraceOrder() (out []string) {
	panic("TODO")
}
