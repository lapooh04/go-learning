package day01

import (
	"errors"
	"math"
	"slices"
	"testing"
)

// These are "table-driven tests" — the dominant Go testing style. A slice of
// anonymous structs holds the cases, and t.Run creates a named subtest for
// each so failures point at the exact case. Read them; you'll write your own
// on Day 6.

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{"plain number", 1, "1"},
		{"divisible by 3", 9, "Fizz"},
		{"divisible by 5", 10, "Buzz"},
		{"divisible by both", 15, "FizzBuzz"},
		{"zero is divisible by both", 0, "FizzBuzz"},
		{"negative plain", -7, "-7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.in); got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{100, "A"}, {90, "A"}, {89, "B"}, {80, "B"},
		{75, "C"}, {60, "D"}, {59, "F"}, {0, "F"},
		{101, "invalid"}, {-1, "invalid"},
	}
	for _, tt := range tests {
		if got := Grade(tt.in); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestSumTo(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{1, 1}, {5, 15}, {100, 5050}, {0, 0}, {-3, 0},
	}
	for _, tt := range tests {
		if got := SumTo(tt.in); got != tt.want {
			t.Errorf("SumTo(%d) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestCollatzSteps(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{1, 0}, {2, 1}, {3, 7}, {6, 8}, {27, 111}, {0, -1}, {-5, -1},
	}
	for _, tt := range tests {
		if got := CollatzSteps(tt.in); got != tt.want {
			t.Errorf("CollatzSteps(%d) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"ascii", "hello", "olleh"},
		{"empty", "", ""},
		{"single", "x", "x"},
		{"palindrome", "racecar", "racecar"},
		// These two fail if you reverse bytes instead of runes.
		{"accented", "héllo", "olléh"},
		{"cjk", "日本語", "語本日"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.in); got != tt.want {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"hello", 2}, {"HELLO", 2}, {"xyz", 0}, {"", 0},
		{"AeIoU", 5}, {"The quick brown fox", 5},
	}
	for _, tt := range tests {
		if got := CountVowels(tt.in); got != tt.want {
			t.Errorf("CountVowels(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		name             string
		in               []int
		wantMin, wantMax int
		wantOK           bool
	}{
		{"typical", []int{3, 1, 4, 1, 5}, 1, 5, true},
		{"single", []int{7}, 7, 7, true},
		{"negatives", []int{-3, -10, -1}, -10, -1, true},
		{"empty", []int{}, 0, 0, false},
		{"nil", nil, 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, gotOK := MinMax(tt.in)
			if gotMin != tt.wantMin || gotMax != tt.wantMax || gotOK != tt.wantOK {
				t.Errorf("MinMax(%v) = (%d, %d, %t), want (%d, %d, %t)",
					tt.in, gotMin, gotMax, gotOK, tt.wantMin, tt.wantMax, tt.wantOK)
			}
		})
	}
}

func TestSafeDivide(t *testing.T) {
	got, err := SafeDivide(10, 2)
	if err != nil {
		t.Fatalf("SafeDivide(10, 2) returned unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("SafeDivide(10, 2) = %d, want 5", got)
	}

	got, err = SafeDivide(10, 0)
	// errors.Is is how you compare against a sentinel — not ==, because the
	// error may be wrapped. More on this Day 3.
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("SafeDivide(10, 0) error = %v, want ErrDivideByZero", err)
	}
	if got != 0 {
		t.Errorf("SafeDivide(10, 0) = %d, want 0 on error", got)
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		in   []float64
		want float64
	}{
		{"typical", []float64{1, 2, 3, 4}, 2.5},
		{"single", []float64{42}, 42},
		{"none", nil, 0},
		{"negatives", []float64{-1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Never compare floats with == — check the difference instead.
			if got := Average(tt.in...); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Average(%v) = %g, want %g", tt.in, got, tt.want)
			}
		})
	}
}

func TestTraceOrder(t *testing.T) {
	want := []string{"start", "end", "three", "two", "one"}
	if got := TraceOrder(); !slices.Equal(got, want) {
		t.Errorf("TraceOrder() = %v, want %v\n"+
			"Remember: defers run last-in-first-out, and they run AFTER the "+
			"return value is set.", got, want)
	}
}
