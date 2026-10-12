// main.go drives a scenario against the simulator.
//
// This file is written FOR you — it's orchestration and printing, not new
// concepts. Treat it as the reward: as you complete each milestone, more of
// this output comes alive.
//
//	go run ./day03/project1
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Guard so an unimplemented skeleton gives you a hint instead of a
	// confusing nil-pointer panic.
	if (Phase(0)).String() == "" {
		fmt.Println("M1 isn't done yet — implement Phase.String() in types.go first.")
		fmt.Println("Run: go test ./day03/project1/ -run TestPhase -v")
		os.Exit(1)
	}

	c := NewCluster()
	s := NewScheduler(c)

	// --- Build a cluster -------------------------------------------------
	section("Cluster bootstrap")

	nodes := []*Node{
		NewNode("worker-1", 4000, 8192),
		NewNode("worker-2", 2000, 4096),
		NewNode("worker-3", 2000, 4096),
	}
	nodes[0].Labels["region"] = "emea"
	nodes[1].Labels["region"] = "emea"
	nodes[2].Labels["region"] = "apac"

	for _, n := range nodes {
		check(c.AddNode(n))
	}
	for _, n := range c.Nodes() {
		fmt.Println(" ", n)
	}

	// --- Create some workloads -------------------------------------------
	section("Creating pods")

	pods := []*Pod{
		NewPod("web", "nginx-1", 500, 512),
		NewPod("web", "nginx-2", 500, 512),
		NewPod("web", "nginx-3", 500, 512),
		NewPod("db", "postgres-1", 1500, 2048),
		NewPod("batch", "report-job", 800, 1024),
	}
	// This one is pinned to the apac region via a node selector.
	pinned := NewPod("edge", "cache-apac", 300, 256)
	pinned.Labels["node.region"] = "apac"
	pods = append(pods, pinned)

	// And this one can never be placed — nothing has 16 cores.
	pods = append(pods, NewPod("ml", "trainer", 16000, 32768))

	for _, p := range pods {
		p.Labels["app"] = p.Namespace
		check(c.AddPod(p))
		fmt.Println(" ", p)
	}

	// --- Schedule ---------------------------------------------------------
	section("Scheduling")

	placed, failed := s.ScheduleAll()
	fmt.Printf("  placed %d pod(s), %d unschedulable\n", placed, len(failed))
	for _, p := range failed {
		fmt.Printf("  ✗ %s could not be placed (requests %s)\n", p.Name(), p.Requests)
	}

	printPlacement(c)

	// --- Inspect ----------------------------------------------------------
	section("Describe")

	if n, err := c.GetNode("worker-1"); err == nil {
		fmt.Println(c.Describe(n))
	}
	if p, err := c.GetPod("db", "postgres-1"); err == nil {
		fmt.Println(c.Describe(p))
	}

	// --- Filter, the kubectl -l way ---------------------------------------
	section(`Pods matching app=web`)

	for _, p := range c.PodsWhere(func(p *Pod) bool { return p.HasLabel("app", "web") }) {
		fmt.Println(" ", p)
	}

	// --- Drain a node -----------------------------------------------------
	section("Draining worker-1")

	evicted, err := s.Drain("worker-1")
	if err != nil {
		fmt.Println("  drain failed:", err)
	} else {
		fmt.Printf("  %d pod(s) could not be rescheduled elsewhere\n", len(evicted))
		for _, p := range evicted {
			fmt.Printf("  ✗ %s is now homeless\n", p.Name())
		}
	}

	printPlacement(c)

	// --- Delete -----------------------------------------------------------
	section("Deleting db/postgres-1")

	if err := c.DeletePod("db", "postgres-1"); err != nil {
		fmt.Println("  delete failed:", err)
	}
	printPlacement(c)
}

// printPlacement renders the cluster as a node-by-node tree.
func printPlacement(c *Cluster) {
	fmt.Println()
	for _, n := range c.Nodes() {
		fmt.Println(" ", n)
		for _, p := range c.PodsOnNode(n.Name()) {
			fmt.Printf("    └─ %s\n", p)
		}
	}
	unscheduled := c.PodsWhere(func(p *Pod) bool { return !p.Scheduled() })
	if len(unscheduled) > 0 {
		fmt.Println("  <unscheduled>")
		for _, p := range unscheduled {
			fmt.Printf("    └─ %s\n", p)
		}
	}
}

func section(title string) {
	fmt.Printf("\n%s\n%s\n", title, strings.Repeat("─", len(title)))
}

// check is a tiny helper for errors that should never happen in this scenario.
//
// Do NOT copy this pattern into real code — exiting on error is only
// acceptable in a demo's main(). Library code returns errors to its caller.
func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal:", err)
		os.Exit(1)
	}
}
