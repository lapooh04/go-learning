package main

import (
	"errors"
	"testing"
)

// Run one milestone at a time:
//
//	go test ./day03/project1/ -run TestPhase      # M1
//	go test ./day03/project1/ -run TestCluster    # M2
//	go test ./day03/project1/ -run TestSchedule   # M3
//	go test ./day03/project1/ -v                  # everything

// ---------------------------------------------------------------------------
// M1 — types
// ---------------------------------------------------------------------------

func TestPhaseString(t *testing.T) {
	tests := []struct {
		in   Phase
		want string
	}{
		{PodPending, "Pending"},
		{PodRunning, "Running"},
		{PodFailed, "Failed"},
		{PodSucceeded, "Succeeded"},
		{Phase(99), "Unknown"},
	}
	for _, tt := range tests {
		if got := tt.in.String(); got != tt.want {
			t.Errorf("Phase(%d).String() = %q, want %q", int(tt.in), got, tt.want)
		}
	}
}

func TestResources(t *testing.T) {
	a := Resources{MilliCPU: 1000, MemoryMi: 512}
	b := Resources{MilliCPU: 250, MemoryMi: 128}

	if got, want := a.Add(b), (Resources{1250, 640}); got != want {
		t.Errorf("Add = %v, want %v", got, want)
	}
	if got, want := a.Sub(b), (Resources{750, 384}); got != want {
		t.Errorf("Sub = %v, want %v", got, want)
	}
	if !a.Fits(b) {
		t.Error("a.Fits(b) = false, want true")
	}
	if b.Fits(a) {
		t.Error("b.Fits(a) = true, want false")
	}
	// Equal resources fit exactly.
	if !a.Fits(a) {
		t.Error("a.Fits(a) = false, want true — exact fit must be allowed")
	}
	// Fits must check BOTH dimensions, not just CPU.
	wide := Resources{MilliCPU: 100, MemoryMi: 99999}
	if a.Fits(wide) {
		t.Error("a.Fits(wide) = true, want false — memory does not fit")
	}
	// Add must not mutate the receiver.
	if a.MilliCPU != 1000 {
		t.Errorf("Add mutated its receiver: a.MilliCPU = %d, want 1000", a.MilliCPU)
	}
}

func TestNodeFree(t *testing.T) {
	n := NewNode("worker-1", 2000, 4096)
	if got, want := n.Free(), (Resources{2000, 4096}); got != want {
		t.Errorf("Free on empty node = %v, want %v", got, want)
	}
	n.Allocated = Resources{MilliCPU: 500, MemoryMi: 1024}
	if got, want := n.Free(), (Resources{1500, 3072}); got != want {
		t.Errorf("Free = %v, want %v", got, want)
	}
}

func TestPodBasics(t *testing.T) {
	p := NewPod("web", "nginx-1", 250, 128)

	if got, want := p.Name(), "web/nginx-1"; got != want {
		t.Errorf("Name = %q, want %q", got, want)
	}
	if p.Scheduled() {
		t.Error("new pod reports Scheduled() = true, want false")
	}
	p.NodeName = "worker-1"
	if !p.Scheduled() {
		t.Error("bound pod reports Scheduled() = false, want true")
	}

	p.Labels["app"] = "web"
	p.Labels["tier"] = ""
	if !p.HasLabel("app", "web") {
		t.Error(`HasLabel("app","web") = false, want true`)
	}
	if p.HasLabel("app", "db") {
		t.Error(`HasLabel("app","db") = true, want false`)
	}
	// Present-but-empty must not be confused with absent.
	if !p.HasLabel("tier", "") {
		t.Error(`HasLabel("tier","") = false, want true — key exists with empty value`)
	}
	if p.HasLabel("missing", "") {
		t.Error(`HasLabel("missing","") = true, want false — key absent`)
	}
}

func TestStringersAreImplemented(t *testing.T) {
	n := NewNode("worker-1", 2000, 4096)
	if n.String() == "" {
		t.Error("Node.String() is empty — M1 not done")
	}
	p := NewPod("web", "nginx-1", 250, 128)
	if p.String() == "" {
		t.Error("Pod.String() is empty — M1 not done")
	}
}

// ---------------------------------------------------------------------------
// M2 — cluster
// ---------------------------------------------------------------------------

func TestClusterAddGet(t *testing.T) {
	c := NewCluster()
	n := NewNode("worker-1", 2000, 4096)

	if err := c.AddNode(n); err != nil {
		t.Fatalf("AddNode: %v", err)
	}
	if err := c.AddNode(n); !errors.Is(err, ErrAlreadyExists) {
		t.Errorf("duplicate AddNode error = %v, want ErrAlreadyExists", err)
	}

	got, err := c.GetNode("worker-1")
	if err != nil {
		t.Fatalf("GetNode: %v", err)
	}
	// Pointer identity: the cluster must hand back the same object, not a copy.
	if got != n {
		t.Error("GetNode returned a different pointer than was added")
	}

	if _, err := c.GetNode("nope"); !errors.Is(err, ErrNotFound) {
		t.Errorf("GetNode(missing) error = %v, want ErrNotFound", err)
	}

	p := NewPod("web", "nginx-1", 250, 128)
	if err := c.AddPod(p); err != nil {
		t.Fatalf("AddPod: %v", err)
	}
	if _, err := c.GetPod("web", "nginx-1"); err != nil {
		t.Errorf("GetPod: %v", err)
	}
	// Same name in a different namespace is a different pod.
	if err := c.AddPod(NewPod("db", "nginx-1", 250, 128)); err != nil {
		t.Errorf("same name in another namespace should be allowed, got %v", err)
	}
}

func TestClusterListsAreSorted(t *testing.T) {
	c := NewCluster()
	for _, name := range []string{"worker-3", "worker-1", "worker-2"} {
		if err := c.AddNode(NewNode(name, 1000, 1024)); err != nil {
			t.Fatal(err)
		}
	}
	// Repeat: map order is randomised per iteration, so an unsorted
	// implementation usually passes once and fails within a few tries.
	for i := 0; i < 20; i++ {
		nodes := c.Nodes()
		if len(nodes) != 3 {
			t.Fatalf("Nodes() returned %d nodes, want 3", len(nodes))
		}
		for j, want := range []string{"worker-1", "worker-2", "worker-3"} {
			if nodes[j].Name() != want {
				t.Fatalf("Nodes()[%d] = %q, want %q (sort by name)",
					j, nodes[j].Name(), want)
			}
		}
	}
}

func TestClusterPodsWhere(t *testing.T) {
	c := NewCluster()
	web := NewPod("web", "nginx-1", 100, 64)
	web.Labels["app"] = "web"
	db := NewPod("db", "postgres-1", 100, 64)
	db.Labels["app"] = "db"
	db.Phase = PodRunning

	for _, p := range []*Pod{web, db} {
		if err := c.AddPod(p); err != nil {
			t.Fatal(err)
		}
	}

	pending := c.PodsWhere(func(p *Pod) bool { return p.Phase == PodPending })
	if len(pending) != 1 || pending[0].Name() != "web/nginx-1" {
		t.Errorf("pending pods = %v, want [web/nginx-1]", pending)
	}

	labelled := c.PodsWhere(func(p *Pod) bool { return p.HasLabel("app", "db") })
	if len(labelled) != 1 || labelled[0].Name() != "db/postgres-1" {
		t.Errorf("app=db pods = %v, want [db/postgres-1]", labelled)
	}

	none := c.PodsWhere(func(p *Pod) bool { return false })
	if len(none) != 0 {
		t.Errorf("predicate matching nothing returned %d pods, want 0", len(none))
	}
}

func TestClusterDeletePodReleasesResources(t *testing.T) {
	c := NewCluster()
	n := NewNode("worker-1", 2000, 4096)
	if err := c.AddNode(n); err != nil {
		t.Fatal(err)
	}
	p := NewPod("web", "nginx-1", 500, 1024)
	if err := c.AddPod(p); err != nil {
		t.Fatal(err)
	}

	s := NewScheduler(c)
	if _, err := s.Schedule(p); err != nil {
		t.Fatalf("Schedule: %v", err)
	}
	if n.Allocated.MilliCPU != 500 {
		t.Fatalf("after scheduling, node Allocated = %v, want 500m", n.Allocated)
	}

	if err := c.DeletePod("web", "nginx-1"); err != nil {
		t.Fatalf("DeletePod: %v", err)
	}
	if n.Allocated != (Resources{0, 0}) {
		t.Errorf("after delete, node Allocated = %v, want 0 — capacity leaked",
			n.Allocated)
	}
	if _, err := c.GetPod("web", "nginx-1"); !errors.Is(err, ErrNotFound) {
		t.Errorf("deleted pod still present, GetPod err = %v", err)
	}
	if err := c.DeletePod("web", "nginx-1"); !errors.Is(err, ErrNotFound) {
		t.Errorf("deleting twice err = %v, want ErrNotFound", err)
	}
}

func TestClusterDescribe(t *testing.T) {
	c := NewCluster()
	n := NewNode("worker-1", 2000, 4096)
	p := NewPod("web", "nginx-1", 250, 128)
	if c.Describe(n) == "" {
		t.Error("Describe(*Node) is empty")
	}
	if c.Describe(p) == "" {
		t.Error("Describe(*Pod) is empty")
	}
	if c.Describe(n) == c.Describe(p) {
		t.Error("Describe returns identical text for a Node and a Pod — " +
			"the type switch is not discriminating")
	}
}

// ---------------------------------------------------------------------------
// M3 — scheduler
// ---------------------------------------------------------------------------

// newTestCluster builds a cluster with two differently-sized nodes.
//
// t.Helper() marks this as a helper, so when a test fails the reported line
// number points at the caller rather than at this function.
func newTestCluster(t *testing.T) (*Cluster, *Scheduler) {
	t.Helper()
	c := NewCluster()
	for _, n := range []*Node{
		NewNode("worker-1", 2000, 4096),
		NewNode("worker-2", 1000, 2048),
	} {
		if err := c.AddNode(n); err != nil {
			t.Fatal(err)
		}
	}
	return c, NewScheduler(c)
}

// mustNode fetches a node or fails the test cleanly.
//
// Without this, an unimplemented GetNode returns (nil, nil) and the next line
// dereferences nil — which panics and takes the whole test binary down with
// it, hiding every other result. Helpers that fail loudly beat helpers that
// crash.
func mustNode(t *testing.T, c *Cluster, name string) *Node {
	t.Helper()
	n, err := c.GetNode(name)
	if err != nil {
		t.Fatalf("GetNode(%q): %v", name, err)
	}
	if n == nil {
		t.Fatalf("GetNode(%q) returned nil — is Cluster.GetNode implemented? (M2)", name)
	}
	return n
}

func TestScheduleFiltersByCapacity(t *testing.T) {
	c, s := newTestCluster(t)

	// Only worker-1 (2000m) can hold this.
	big := NewPod("web", "big", 1500, 1024)
	if err := c.AddPod(big); err != nil {
		t.Fatal(err)
	}

	candidates := s.Filter(big)
	if len(candidates) != 1 || candidates[0].Name() != "worker-1" {
		t.Fatalf("Filter returned %v, want [worker-1]", candidates)
	}

	n, err := s.Schedule(big)
	if err != nil {
		t.Fatalf("Schedule: %v", err)
	}
	if n == nil {
		t.Fatal("Schedule returned (nil, nil) — is Scheduler.Schedule implemented? (M3)")
	}
	if n.Name() != "worker-1" {
		t.Errorf("scheduled onto %s, want worker-1", n.Name())
	}
	if big.Phase != PodRunning {
		t.Errorf("pod phase = %v, want Running", big.Phase)
	}
	if big.NodeName != "worker-1" {
		t.Errorf("pod NodeName = %q, want worker-1", big.NodeName)
	}
	if n.Allocated.MilliCPU != 1500 {
		t.Errorf("node Allocated = %v, want 1500m charged", n.Allocated)
	}
}

func TestScheduleSkipsCordonedNodes(t *testing.T) {
	c, s := newTestCluster(t)
	w1 := mustNode(t, c, "worker-1")
	w1.Schedulable = false

	p := NewPod("web", "p1", 1500, 1024) // only fits worker-1
	if err := c.AddPod(p); err != nil {
		t.Fatal(err)
	}

	if _, err := s.Schedule(p); !errors.Is(err, ErrUnschedulable) {
		t.Errorf("Schedule onto cordoned-only cluster err = %v, want ErrUnschedulable", err)
	}
}

func TestScheduleSpreadsLoad(t *testing.T) {
	c, s := newTestCluster(t)

	// worker-1 has 2000m/4096Mi, worker-2 has 1000m/2048Mi. Both fit a small
	// pod, but worker-1 is proportionally emptier after placement, so
	// least-allocated scoring must pick it.
	p := NewPod("web", "small", 100, 128)
	if err := c.AddPod(p); err != nil {
		t.Fatal(err)
	}
	n, err := s.Schedule(p)
	if err != nil {
		t.Fatalf("Schedule: %v", err)
	}
	if n == nil {
		t.Fatal("Schedule returned (nil, nil) — is Scheduler.Schedule implemented? (M3)")
	}
	if n.Name() != "worker-1" {
		t.Errorf("scheduled onto %s, want worker-1 (the least-allocated node)", n.Name())
	}
}

func TestBindRejectsOvercommit(t *testing.T) {
	c, s := newTestCluster(t)
	w2 := mustNode(t, c, "worker-2") // 1000m / 2048Mi

	p := NewPod("web", "huge", 5000, 128)
	if err := c.AddPod(p); err != nil {
		t.Fatal(err)
	}

	err := s.Bind(p, w2)
	if err == nil {
		t.Fatal("Bind overcommitted the node without error")
	}
	// errors.As digs a typed error out of a possibly-wrapped chain.
	var ire *InsufficientResourcesError
	if !errors.As(err, &ire) {
		t.Fatalf("Bind error = %v (%T), want *InsufficientResourcesError", err, err)
	}
	if ire.NodeName != "worker-2" {
		t.Errorf("error names node %q, want worker-2", ire.NodeName)
	}
}

func TestNodeSelector(t *testing.T) {
	n := NewNode("worker-1", 2000, 4096)
	n.Labels["region"] = "emea"

	plain := NewPod("web", "plain", 100, 64)
	if !NodeSelector(n, plain) {
		t.Error("pod with no selector should match any node")
	}

	matching := NewPod("web", "matching", 100, 64)
	matching.Labels["node.region"] = "emea"
	if !NodeSelector(n, matching) {
		t.Error("node.region=emea should match a node labelled region=emea")
	}

	mismatched := NewPod("web", "mismatched", 100, 64)
	mismatched.Labels["node.region"] = "apac"
	if NodeSelector(n, mismatched) {
		t.Error("node.region=apac should NOT match a node labelled region=emea")
	}

	// A normal label must not be treated as a selector.
	normal := NewPod("web", "normal", 100, 64)
	normal.Labels["app"] = "web"
	if !NodeSelector(n, normal) {
		t.Error("a non-'node.' label must not act as a selector")
	}
}

func TestScheduleAllReportsFailures(t *testing.T) {
	c, s := newTestCluster(t) // 3000m total across both nodes

	if err := c.AddPod(NewPod("web", "a", 800, 512)); err != nil {
		t.Fatal(err)
	}
	if err := c.AddPod(NewPod("web", "b", 800, 512)); err != nil {
		t.Fatal(err)
	}
	if err := c.AddPod(NewPod("web", "impossible", 9000, 512)); err != nil {
		t.Fatal(err)
	}

	placed, failed := s.ScheduleAll()
	if placed != 2 {
		t.Errorf("placed = %d, want 2", placed)
	}
	if len(failed) != 1 || failed[0].Name() != "web/impossible" {
		t.Errorf("failed = %v, want [web/impossible] — one bad pod must not "+
			"stop the others", failed)
	}
}

func TestDrainReschedules(t *testing.T) {
	c, s := newTestCluster(t)

	p := NewPod("web", "mover", 400, 256)
	if err := c.AddPod(p); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Schedule(p); err != nil {
		t.Fatal(err)
	}
	original := p.NodeName

	evicted, err := s.Drain(original)
	if err != nil {
		t.Fatalf("Drain: %v", err)
	}
	if len(evicted) != 0 {
		t.Errorf("evicted (unplaceable) = %v, want none — the other node had room", evicted)
	}

	drained := mustNode(t, c, original)
	if drained.Schedulable {
		t.Error("drained node is still Schedulable")
	}
	if drained.Allocated != (Resources{0, 0}) {
		t.Errorf("drained node still has Allocated = %v, want 0", drained.Allocated)
	}
	if p.NodeName == original {
		t.Errorf("pod is still on the drained node %q", original)
	}
	if p.Phase != PodRunning {
		t.Errorf("rescheduled pod phase = %v, want Running", p.Phase)
	}

	if _, err := s.Drain("ghost"); !errors.Is(err, ErrNotFound) {
		t.Errorf("Drain(missing) err = %v, want ErrNotFound", err)
	}
}
