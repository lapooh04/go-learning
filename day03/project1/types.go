// Package main implements clustersim, a miniature cluster control plane.
//
// types.go holds the domain model: what a Node is, what a Pod is, and the
// interface that lets the rest of the program treat them uniformly.
package main

import "fmt"

// ---------------------------------------------------------------------------
// Resources
// ---------------------------------------------------------------------------

// Resources is a CPU/memory pair. CPU is in millicores (1 core = 1000m),
// memory in MiB. Keeping them as plain ints avoids quantity parsing entirely.
//
// Note this is a value type with no pointer fields, so copying it is cheap and
// safe — that's why the methods below use value receivers.
type Resources struct {
	MilliCPU int
	MemoryMi int
}

// Add returns the sum of r and o.
//
// It returns a NEW Resources rather than mutating r. With a value receiver,
// mutating r would only change the local copy anyway — a classic Go beginner
// trap worth proving to yourself in the debugger.
func (r Resources) Add(o Resources) Resources {
	// TODO (M1)
	return Resources{}
}

// Sub returns r minus o. Values may go negative; the caller decides if that
// matters.
func (r Resources) Sub(o Resources) Resources {
	// TODO (M1)
	return Resources{}
}

// Fits reports whether o can fit within r (i.e. r has at least as much of both
// CPU and memory as o needs).
func (r Resources) Fits(o Resources) bool {
	// TODO (M1)
	return false
}

// String makes Resources satisfy fmt.Stringer, so it prints as "500m/256Mi"
// instead of "{500 256}" anywhere you use %v or Println.
func (r Resources) String() string {
	return fmt.Sprintf("%dm/%dMi", r.MilliCPU, r.MemoryMi)
}

// ---------------------------------------------------------------------------
// Phase — an enum, Go style
// ---------------------------------------------------------------------------

// Phase is a pod's lifecycle state.
//
// Go has no enum keyword. The idiom is a named integer type plus a const block
// using iota, which auto-increments: PodPending is 0, PodRunning is 1, and so
// on. Defining it as its own type (not a bare int) means the compiler stops
// you passing an arbitrary int where a Phase is expected.
type Phase int

const (
	PodPending Phase = iota
	PodRunning
	PodFailed
	PodSucceeded
)

// String makes Phase printable. Without it, fmt prints the raw number.
//
// Return "Pending", "Running", "Failed", "Succeeded", or "Unknown" for any
// value outside the defined set.
//
// Hint: a switch on the receiver. Consider also a package-level
// [...]string array indexed by the phase — both are idiomatic; try the switch
// first since it handles the out-of-range case more naturally.
func (p Phase) String() string {
	// TODO (M1)
	return ""
}

// ---------------------------------------------------------------------------
// The Object interface
// ---------------------------------------------------------------------------

// Object is anything the cluster can store. Both *Node and *Pod implement it.
//
// This is the heart of the exercise. Note how small it is — three methods.
// Go interfaces describe the minimum behaviour a caller needs, not everything
// a type can do. Compare io.Reader: one method, and it's the most reused
// interface in the entire language.
//
// Also note: nothing anywhere says "Node implements Object". Satisfaction is
// implicit and checked at compile time wherever you actually use a *Node as an
// Object.
type Object interface {
	// Kind returns "Node" or "Pod" — the same idea as a Kubernetes kind.
	Kind() string
	// Name returns the object's name, unique within its kind.
	Name() string
	fmt.Stringer // embedding an interface: Object also requires String() string
}

// Compile-time assertions that our types really do satisfy Object.
//
// This is a common Go trick: assign a nil pointer to the interface type and
// discard it. It costs nothing at runtime, but if *Node ever stops satisfying
// Object you get an error HERE, with a clear message, instead of a confusing
// one at some distant call site.
var (
	_ Object = (*Node)(nil)
	_ Object = (*Pod)(nil)
)

// ---------------------------------------------------------------------------
// Node
// ---------------------------------------------------------------------------

// Node is a machine that can run pods.
type Node struct {
	NodeName    string
	Labels      map[string]string
	Capacity    Resources // total resources on the box
	Allocated   Resources // sum of the requests of pods bound here
	Schedulable bool      // false once cordoned/drained
}

// NewNode builds a Node with an initialised Labels map.
//
// Constructor functions like this are the Go answer to constructors. The
// reason this one matters: the zero value of a map is nil, and writing to a
// nil map PANICS. Reading from one is fine and returns the zero value, which
// is exactly the sort of asymmetry that produces a 2am incident. Initialise
// maps when you build the struct.
func NewNode(name string, milliCPU, memoryMi int) *Node {
	return &Node{
		NodeName:    name,
		Labels:      make(map[string]string),
		Capacity:    Resources{MilliCPU: milliCPU, MemoryMi: memoryMi},
		Schedulable: true,
	}
}

func (n *Node) Kind() string { return "Node" }
func (n *Node) Name() string { return n.NodeName }

// Free returns the resources still available on n.
func (n *Node) Free() Resources {
	// TODO (M1)
	return Resources{}
}

// String renders the node roughly like a `kubectl get nodes` row, e.g.
//
//	Node/worker-1   Ready        1500m/3072Mi free
//	Node/worker-2   Cordoned      500m/1024Mi free
//
// Exact format is yours to choose — just make it readable and include the
// name, whether it's schedulable, and how much is free.
func (n *Node) String() string {
	// TODO (M1)
	return ""
}

// ---------------------------------------------------------------------------
// Pod
// ---------------------------------------------------------------------------

// Pod is a unit of work that must be placed on a Node.
type Pod struct {
	PodName   string
	Namespace string
	Labels    map[string]string
	Requests  Resources
	NodeName  string // empty until scheduled
	Phase     Phase
}

// NewPod builds an unscheduled, Pending pod.
func NewPod(namespace, name string, milliCPU, memoryMi int) *Pod {
	return &Pod{
		PodName:   name,
		Namespace: namespace,
		Labels:    make(map[string]string),
		Requests:  Resources{MilliCPU: milliCPU, MemoryMi: memoryMi},
		Phase:     PodPending,
	}
}

func (p *Pod) Kind() string { return "Pod" }

// Name returns the namespaced name, "namespace/name", which is what makes a
// pod unique across the cluster.
func (p *Pod) Name() string {
	// TODO (M1)
	return ""
}

// Scheduled reports whether p has been bound to a node.
func (p *Pod) Scheduled() bool {
	// TODO (M1)
	return false
}

// String renders the pod roughly like a `kubectl get pods` row, e.g.
//
//	Pod/web/nginx-1   Running   on worker-1   250m/128Mi
//	Pod/web/nginx-9   Pending   unscheduled   250m/128Mi
func (p *Pod) String() string {
	// TODO (M1)
	return ""
}

// HasLabel reports whether p carries the given label key=value pair.
//
// Hint: this is the comma-ok idiom. `v, ok := p.Labels[key]` tells you both
// the value and whether the key was present at all — you need the second
// return to distinguish "missing" from "present but empty".
func (p *Pod) HasLabel(key, value string) bool {
	// TODO (M1)
	return false
}
