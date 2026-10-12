# Day 1 — Syntax, types, control flow

Goal by tonight: you can read any simple Go file without stopping, and the
compiler's error messages make sense to you.

## Phase 1 — Concepts (~60 min)

Read [A Tour of Go](https://go.dev/tour/), sections **Basics** and
**Flow control**. Use the editable playground — change the examples and break
them on purpose.

Things worth pausing on, because they differ from most languages:

**No implicit numeric conversion.** `var i int = 42; var f float64 = i` does
not compile. You must write `float64(i)`. Go would rather be noisy than guess.

**Zero values, not null.** Every declared variable has a usable zero value:
`0` for numbers, `""` for strings, `false` for bools, `nil` for pointers,
slices, maps, channels, and interfaces. There is no "uninitialized".

**`:=` only works inside functions**, and only when at least one variable on
the left is new. At package level you must use `var` or `const`.

**Exported means capitalized.** `FizzBuzz` is visible to other packages;
`fizzbuzz` is not. That's the entire access-control system.

**`for` is the only loop keyword.** Four forms:
```go
for i := 0; i < n; i++ {}   // classic
for condition {}            // while
for {}                      // infinite
for i, v := range xs {}     // range
```

**`switch` doesn't fall through** and cases needn't be constants. A
conditionless `switch { case x > 10: ... }` is the idiomatic if/else ladder.

**`if` can carry an init statement:** `if v, err := f(); err != nil {`. This
scopes `v` and `err` to the if/else, which is why you'll see it constantly.

**`defer` runs when the function returns**, LIFO order, even on panic. It's how
Go does `finally` — mostly for closing files and unlocking mutexes.

## Phase 2 — Drills (~90 min)

```bash
cd ~/go-learning
go test ./day01/            # all failing right now
go test ./day01/ -run Fizz  # just one, while you work on it
go test ./day01/ -v         # see every subtest
```

Open `exercises.go` and work down from the top. Ten functions, each with a doc
comment telling you what it must do and a hint about the Go-specific part.
Don't look up solutions — if you're stuck for more than ten minutes on one, ask
me and I'll give you a nudge rather than the answer.

Two of them are traps worth knowing about in advance:
- `ReverseString` has tests with accented and CJK characters. A byte-wise
  reverse passes the ASCII cases and corrupts the rest.
- `TraceOrder` tests that you understand defer ordering *and* that deferred
  closures can mutate a named return value.

## Phase 3 — Project (~120 min): `gowc`

Build a `wc` clone at `day01/cmd/gowc/main.go`. The skeleton is there with the
requirements as TODOs.

```bash
go run ./day01/cmd/gowc -l -w README.md
cat README.md | go run ./day01/cmd/gowc
go build -o /tmp/gowc ./day01/cmd/gowc && /tmp/gowc -c go.mod
```

This is deliberately an excuse to touch `os`, `bufio`, `flag`, and error
handling before Day 4 formalizes them. Rough edges are fine today.

## Phase 4 — Review (~30 min)

```bash
gofmt -l .      # must print nothing — Go has exactly one formatting style
go vet ./...    # catches real mistakes the compiler permits
go doc strings  # skim what's in the strings package
```

Then read the source of something you used:
```bash
go doc -src strings.Count
```
Reading stdlib source is the single fastest way to absorb Go idiom. It is
written to be read.

## Done when

- `go test ./day01/` is green
- `gowc` handles a file argument, multiple files, and stdin
- You can explain, out loud, why `ReverseString` needs `[]rune`
