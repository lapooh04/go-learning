# clustersim — a mini OpenShift/Kubernetes cluster simulator

Using the concepts learnt so far (loops, if/else, switch, pointers, functions,
methods, structs, maps, arrays, slices, interfaces) to build a miniature
cluster control plane that you can actually drive from the command line.

No external dependencies, no real cluster, no YAML yet. Pure stdlib.

## Why this project

You already know what a Node, a Pod, and a scheduler *are* from OpenShift. That
means you can spend 100% of your attention on the Go, not on the domain. And
the domain maps onto Go concepts almost suspiciously well:

| Cluster concept | Go concept it teaches |
|---|---|
| Node and Pod are both API objects | **interfaces** — one `Object` type for both |
| Pod phase (Pending/Running/Failed) | **iota enums + Stringer** |
| Cluster holds objects by name | **maps**, and the comma-ok idiom |
| `kubectl get pods -l app=web` | **slices + closures** as filter predicates |
| Scheduling mutates a Node's allocation | **pointer receivers** vs value receivers |
| "node not found", "insufficient CPU" | **sentinel + custom error types** |
| `describe` prints any object | **type switch** |
| The scheduler itself | an actual algorithm worth writing |

## Milestones

Do these in order. Each is independently runnable, so you always have something
that works.

**M1 — Types and Stringer** (~45 min)
Fill in `types.go`: the `Object` interface, `Phase` enum with its `String()`
method, and `String()` for `Node` and `Pod`. Get `go test ./day03/project1/ -run TestPhase` green.

**M2 — Cluster state** (~60 min)
Fill in `cluster.go`: add, get, delete, and list objects. This is where you
meet the comma-ok idiom and sentinel errors properly. Target `-run TestCluster`.

**M3 — The scheduler** (~60 min)
Fill in `scheduler.go`. Filter nodes that can fit the pod, score the survivors,
bind the pod to the winner. This is the genuinely interesting part, and it's a
real (simplified) version of what kube-scheduler does. Target `-run TestSchedule`.

**M4 — Drive it** (~30 min)
`main.go` runs a scenario: build a cluster, schedule pods, drain a node, watch
pods get rescheduled. Make the output readable.

**Stretch — only after M1–M4 are green**
Simulate a node failing and reconcile the cluster back to the desired state in
a goroutine with a `time.Ticker`. That's your bridge into Day 5 concurrency.

## Running it

```bash
cd ~/go-learning
go test ./day03/project1/           # the drills
go test ./day03/project1/ -v -run TestSchedule
go run ./day03/project1             # the scenario
```

## Rules

Same as the drills: the skeleton tells you *what*, never *how*. If you're stuck
more than ten minutes, ask for a nudge.

Resources are deliberately simple integers — CPU in **millicores** (1 core =
1000m) and memory in **MiB**. No quantity parsing. That's a Day 4 problem.
