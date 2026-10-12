# Learning Go

My notes and practice code from learning Go from scratch, starting October 2026.

Coming from another language, so the focus here is on what Go does
*differently* — value semantics, implicit interfaces, explicit errors, and
goroutines — rather than general programming basics.

Most files are heavily commented. The comments are the notes; the code is the
proof I understood them.

## Layout

Each `pracN.go` is a standalone program exploring one concept from
[A Tour of Go](https://go.dev/tour/). The mini projects put several concepts
together into something that actually runs.

```
day01/  variables, types, zero values, type conversion, constants
day02/  for loops, if/else, switch, defer, pointers, structs, slices, maps
day03/  maps exercise, closures, methods, interfaces, type assertions
day04/  Stringer, errors, io.Reader, goroutines & channels, WaitGroups, generics
day05-mini-projects/
day06/
day07/
```

### Mini projects

| Project | What it does | Go concepts |
|---|---|---|
| `day05-mini-projects/miniproj1.go` | Cloud budget estimator — takes CPU/RAM/storage, prints a formatted receipt | constants, user input, formatted output |
| `day05-mini-projects/miniproj2.go` | GitHub user dashboard — fetches the GitHub API and renders a terminal profile | `net/http`, `encoding/json`, structs, `defer` for closing the response body |
| `day05-mini-projects/miniproj3.go` | CLI typing test — random prompt, measures WPM and accuracy | `bufio` on stdin, `math/rand`, `time`, `strings` |
| `day06/miniproj4.go` | URL shortener backend — POST a long URL, get a short code, follow the redirect | HTTP handlers, JSON decoding, maps as an in-memory store, HTTP redirects |
| `day06/frontend/` | Static file server + HTML page for the shortener | `http.FileServer`, serving static assets |
| `day07/miniproj5.go` | CodeClock CLI - Measures the time spent on a task, saves session data in JSON format.

## Running things

Each `pracN.go` and mini project is its own `package main` with its own
`main()`, so they are run one file at a time rather than as a single module:

```bash
go run day02/prac9.go
go run day05-mini-projects/miniproj3.go
```

`go build ./...` across the whole repo will **not** work, and that's expected —
several directories contain multiple `main()` functions, which Go does not
allow in one package. This is a notebook, not an application.
whi
The URL shortener needs two terminals:

```bash
go run day06/miniproj4.go          # backend
cd day06/frontend && go run frontend_launcher.go   # then open localhost:3000
```

## Setup

Go 1.27.1. The `.vscode/` directory has working settings for the Go extension,
including explicit toolchain paths — necessary because VS Code is installed as
a Flatpak and does not inherit `PATH` from the shell.

## Status

Still in progress. `day01/exercises.go` and `day03/project1/` are scaffolded
exercise sets I set aside in favour of working through the Tour directly and
building the mini projects instead.
