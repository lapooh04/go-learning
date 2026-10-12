// Command gowc is a small clone of the Unix `wc` utility.
//
// This is Day 1's project. The structure is here; the logic is yours.
//
// Target behaviour:
//
//	gowc -l -w file.txt      count lines and words in file.txt
//	gowc file1.txt file2.txt count both, then print a "total" line
//	cat x | gowc             read stdin when no files are given
//	gowc                     with no flags, show lines, words, and bytes
//
// Exit 0 on success, 1 if any file could not be read (but still process the
// rest — don't bail on the first error).
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// counts holds the tally for one input source.
type counts struct {
	lines int
	words int
	bytes int
}

func main() {
	// flag.Bool returns a *bool. Dereference with *showLines to read it.
	showLines := flag.Bool("l", false, "count lines")
	showWords := flag.Bool("w", false, "count words")
	showBytes := flag.Bool("c", false, "count bytes")
	flag.Parse()

	// Go treats an unused local variable as a compile ERROR, not a warning.
	// This line keeps the skeleton compiling; delete it once you actually use
	// the flags. `_` is the blank identifier — assign to it to discard a value.
	_, _, _ = showLines, showWords, showBytes

	// TODO: if none of the three flags were set, default to all three on.

	paths := flag.Args() // non-flag arguments

	if len(paths) == 0 {
		// TODO: count os.Stdin and print the result with no filename.
		// os.Stdin is an *os.File, which satisfies io.Reader.
		_ = os.Stdin
		return
	}

	// TODO: for each path: open it, count it, print its row, and accumulate
	// into a running total. Remember `defer f.Close()` — but note that a defer
	// inside a loop doesn't run until the function ends, so a helper function
	// per file is the cleaner shape here.
	//
	// TODO: if len(paths) > 1, print a final "total" row.
	//
	// TODO: exit 1 if any file failed. Print errors to os.Stderr, not stdout.
	for _, p := range paths {
		_ = p
	}
}

// count reads everything from r and returns the tallies.
//
// Taking an io.Reader rather than a filename is what makes this testable and
// lets the same function handle both files and stdin. Getting used to writing
// against io.Reader instead of *os.File is one of the most valuable habits in
// Go — you'll meet the reasoning properly on Day 3.
//
// Hint: bufio.Scanner with bufio.ScanWords / bufio.ScanLines is one way, but
// you can only scan a reader once — so consider counting manually, or use
// bufio.Reader and walk it yourself.
func count(r io.Reader) (counts, error) {
	// TODO
	return counts{}, fmt.Errorf("not implemented")
}

// format renders c according to which flags are enabled, followed by label.
// Real wc right-aligns its numbers; matching that exactly is a nice stretch.
func format(c counts, label string, lines, words, bytes bool) string {
	// TODO
	return ""
}

// countFile opens path, counts it, and closes it.
func countFile(path string) (counts, error) {
	// TODO: os.Open, defer Close, call count.
	// Wrap any error with context: fmt.Errorf("gowc: %s: %w", path, err)
	return counts{}, fmt.Errorf("not implemented")
}
