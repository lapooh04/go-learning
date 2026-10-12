// scheduler.go is the interesting part: deciding which node a pod lands on.
//
// This is a genuinely simplified kube-scheduler. The real one runs the same
// three phases — filter ("predicates"), score ("priorities"), bind — just with
// a lot more plugins in each.
package main

import "errors"

// ErrUnschedulable means no node in the cluster could accept the pod.
var ErrUnschedulable = errors.New("no suitable node for pod")

// Scheduler places pods onto nodes.
//
// It's a struct with one field rather than a bare function so you can hang
// configuration off it later (a scoring strategy, a node-selector policy).
// That's a common Go shape for a component with behaviour and settings.
type Scheduler struct {
	Cluster *Cluster
}

func NewScheduler(c *Cluster) *Scheduler {
	return &Scheduler{Cluster: c}
}

// Filter returns the nodes that could run p.
//
// A node survives the filter when ALL of these hold:
//   - it is Schedulable (not cordoned)
//   - its Free() resources Fit the pod's Requests
//   - it satisfies the pod's node selector (see NodeSelector below)
//
// Return an empty slice, not nil, when nothing fits — though note that in Go
// both have len() == 0 and both range safely, so callers rarely care. Knowing
// that nil slices are usable is itself worth internalising.
func (s *Scheduler) Filter(p *Pod) []*Node {
	// TODO (M3)
	return nil
}

// Score ranks a candidate node for pod p. Higher is better.
//
// Implement "least allocated wins": the node left with the highest percentage
// of free resources after placing the pod scores highest. This spreads load
// instead of packing one node full, and it's roughly what the real
// LeastAllocated priority does.
//
// A workable formula, averaging the two dimensions:
//
//	cpuFree := (node.Capacity.MilliCPU - node.Allocated.MilliCPU - p.Requests.MilliCPU)
//	cpuPct  := 100 * cpuFree / node.Capacity.MilliCPU
//	... same for memory, then return the average
//
// Watch out: integer division truncates, and a Capacity of 0 divides by zero.
// Decide what a zero-capacity node should score.
func (s *Scheduler) Score(n *Node, p *Pod) int {
	// TODO (M3)
	return 0
}

// Bind attaches p to n and charges the pod's requests against the node.
//
// Three things must happen together:
//  1. p.NodeName = n.Name()
//  2. p.Phase = PodRunning
//  3. n.Allocated grows by p.Requests
//
// If you forget (3), the scheduler will happily overcommit every node — a
// good bug to cause once on purpose and watch the tests catch.
//
// Return an *InsufficientResourcesError if n cannot actually fit p. Bind is
// the last line of defence; don't trust that Filter was called first.
func (s *Scheduler) Bind(p *Pod, n *Node) error {
	// TODO (M3)
	return nil
}

// Schedule finds the best node for p and binds it.
//
// Filter, then pick the highest-scoring survivor, then Bind. Return
// ErrUnschedulable if the filter came back empty.
//
// Tie-break deterministically — on node name, say. Otherwise the result
// depends on map iteration order and your tests will flake intermittently,
// which is far more annoying to debug than failing outright.
func (s *Scheduler) Schedule(p *Pod) (*Node, error) {
	// TODO (M3)
	return nil, nil
}

// ScheduleAll schedules every Pending pod in the cluster.
//
// Return the count placed and a slice of the pods that could not be placed.
// Don't stop at the first failure: one unschedulable pod must not block the
// rest. Partial success with a report of what failed is the normal shape for
// this kind of batch operation.
func (s *Scheduler) ScheduleAll() (placed int, failed []*Pod) {
	// TODO (M3)
	return 0, nil
}

// ---------------------------------------------------------------------------
// Node selectors
// ---------------------------------------------------------------------------

// NodeSelector reports whether node n satisfies pod p's node selector.
//
// Convention for this project: any pod label whose key starts with "node." is
// a selector, matched against the node's labels with the prefix stripped. So a
// pod labelled "node.region=emea" only lands on nodes labelled "region=emea".
// A pod with no "node." labels matches every node.
//
// Hint: strings.CutPrefix returns the remainder and a bool, which is a tidy
// fit for this loop.
func NodeSelector(n *Node, p *Pod) bool {
	// TODO (M3)
	return true
}

// ---------------------------------------------------------------------------
// Drain
// ---------------------------------------------------------------------------

// Drain cordons a node and reschedules its pods elsewhere.
//
// Order matters. Mark the node unschedulable FIRST, or the scheduler may
// cheerfully place the evicted pods straight back onto the node you're
// draining. Then for each pod on it: release its resources, reset it to
// Pending and unbound, and schedule it again.
//
// Return the pods that could not be placed anywhere else.
//
// Another trap: you are mutating the set of pods on the node while iterating
// over it. Take the slice from PodsOnNode first and iterate that snapshot.
func (s *Scheduler) Drain(nodeName string) (evicted []*Pod, err error) {
	// TODO (M3)
	return nil, nil
}
