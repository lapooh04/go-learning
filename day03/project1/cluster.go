// cluster.go holds the cluster state and the "API server" operations on it:
// add, get, list, delete. This is where maps and errors get a real workout.
package main

import (
	"errors"
	"fmt"
)

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

// Sentinel errors: package-level values callers compare with errors.Is.
//
// Why errors.Is and not ==? Because the error you receive may be WRAPPED —
// fmt.Errorf("...: %w", ErrNotFound) produces a new error that is not == to
// ErrNotFound but still "is" it. errors.Is unwraps the chain for you.
var (
	ErrNotFound      = errors.New("object not found")
	ErrAlreadyExists = errors.New("object already exists")
)

// InsufficientResourcesError is a CUSTOM error type, as opposed to a sentinel.
//
// Use one of these when the caller might want structured detail about the
// failure rather than just its identity — here, which node failed and by how
// much it fell short. Callers dig it out with errors.As.
type InsufficientResourcesError struct {
	NodeName  string
	Requested Resources
	Free      Resources
}

// Error makes the type satisfy the built-in error interface. Note the pointer
// receiver: by convention you then return &InsufficientResourcesError{...},
// and errors.As targets **InsufficientResourcesError.
func (e *InsufficientResourcesError) Error() string {
	return fmt.Sprintf("node %s cannot fit request %s (free: %s)",
		e.NodeName, e.Requested, e.Free)
}

// ---------------------------------------------------------------------------
// Cluster
// ---------------------------------------------------------------------------

// Cluster is the whole simulated control plane: an in-memory store of objects.
//
// Nodes are keyed by name, pods by their namespaced "ns/name". Storing
// POINTERS matters — you want Get to hand back a reference you can mutate,
// not a copy. Try switching to map[string]Node later and watch scheduling
// silently stop working; it's an instructive five minutes.
type Cluster struct {
	nodes map[string]*Node
	pods  map[string]*Pod
}

// NewCluster returns an empty, ready-to-use Cluster.
func NewCluster() *Cluster {
	return &Cluster{
		nodes: make(map[string]*Node),
		pods:  make(map[string]*Pod),
	}
}

// AddNode stores n, or returns ErrAlreadyExists if a node of that name is
// already present.
//
// Hint: wrap the sentinel with context so the message is useful but
// errors.Is still works:
//
//	return fmt.Errorf("add node %s: %w", n.Name(), ErrAlreadyExists)
func (c *Cluster) AddNode(n *Node) error {
	// TODO (M2)
	return nil
}

// AddPod stores p, or returns ErrAlreadyExists.
func (c *Cluster) AddPod(p *Pod) error {
	// TODO (M2)
	return nil
}

// GetNode returns the node called name, or ErrNotFound.
func (c *Cluster) GetNode(name string) (*Node, error) {
	// TODO (M2)
	return nil, nil
}

// GetPod returns the pod called "namespace/name", or ErrNotFound.
func (c *Cluster) GetPod(namespace, name string) (*Pod, error) {
	// TODO (M2)
	return nil, nil
}

// Nodes returns all nodes, sorted by name.
//
// Sorting is not cosmetic here: Go deliberately RANDOMISES map iteration
// order, so an unsorted version returns a different order on every call and
// your tests will flake. Finding that out the hard way is a rite of passage;
// consider this your warning.
//
// Hint: the `slices` and `sort` packages are both in the stdlib.
// slices.SortFunc with a comparison on Name() is the modern way.
func (c *Cluster) Nodes() []*Node {
	// TODO (M2)
	return nil
}

// Pods returns all pods, sorted by namespaced name.
func (c *Cluster) Pods() []*Pod {
	// TODO (M2)
	return nil
}

// PodsWhere returns every pod for which pred returns true, sorted by name.
//
// Taking a function as a parameter is what makes this one function cover
// `get pods -l app=web`, `get pods --field-selector status.phase=Pending`,
// and anything else you think of later. Callers pass a closure:
//
//	pending := c.PodsWhere(func(p *Pod) bool { return p.Phase == PodPending })
//	web     := c.PodsWhere(func(p *Pod) bool { return p.HasLabel("app", "web") })
//
// Hint: you already wrote closures in prac13.go. Same idea, other direction —
// there you returned functions, here you receive one.
func (c *Cluster) PodsWhere(pred func(*Pod) bool) []*Pod {
	// TODO (M2)
	return nil
}

// PodsOnNode returns the pods currently bound to nodeName.
//
// Hint: this should be one line that calls PodsWhere.
func (c *Cluster) PodsOnNode(nodeName string) []*Pod {
	// TODO (M2)
	return nil
}

// DeletePod removes a pod and releases its resources back to its node.
//
// Careful: if the pod was scheduled you must also subtract its requests from
// the node's Allocated, or the cluster slowly leaks capacity. This is exactly
// the kind of bug that pointer-vs-value confusion causes, so it's worth doing
// deliberately.
func (c *Cluster) DeletePod(namespace, name string) error {
	// TODO (M2)
	return nil
}

// Describe renders any Object with kind-specific detail.
//
// This is what the Object interface buys you: one function that accepts both
// *Node and *Pod. Use a TYPE SWITCH (as in your prac16.go) to recover the
// concrete type and print the extra fields:
//
//	switch o := obj.(type) {
//	case *Node: // o is a *Node here — print capacity, allocated, pod count
//	case *Pod:  // o is a *Pod here  — print namespace, requests, node
//	}
//
// Include a default case. Someone will add a third kind eventually.
func (c *Cluster) Describe(obj Object) string {
	// TODO (M2)
	return ""
}
