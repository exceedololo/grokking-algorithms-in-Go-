package main

import (
	"fmt"
	"math"
)

var graph map[string]map[string]float64
var costs map[string]float64
var parents map[string]string
var processed map[string]bool

func main() {
	// the graph
	graph = make(map[string]map[string]float64)
	graph["start"] = make(map[string]float64)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]float64)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]float64)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]float64)

	// the costs table
	costs = make(map[string]float64)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.Inf(1)

	// the parents table
	parents = make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	// initialize the processed map
	processed = make(map[string]bool)

	// Find the lowest-cost node that you haven't processed yet.
	node := find_lowest_cost_node(costs)
	// If you've processed all the nodes, this while loop is done.

	for node != "" {
		cost := costs[node]
		// Go through all the neighbors of this node.
		neighbors := graph[node]

		for node, _ := range neighbors {
			new_cost := cost + neighbors[node]
			// If it's cheaper to get to this neighbor by going through this node...
			if costs[node] > new_cost {
				// ... update the cost for this node.
				costs[node] = new_cost
				// This node becomes the new parent for this neighbor.
				parents[node] = node
			}
		}
		// Mark the node as processed.
		processed[node] = true
		// Find the next node to process, and loop.
		node = find_lowest_cost_node(costs)
	}

	fmt.Println(costs)
}

func find_lowest_cost_node(costs map[string]float64) string {
	lowest_cost := math.Inf(1)
	lowest_cost_node := ""

	for node, _ := range costs {
		cost := costs[node]
		// If it's the lowest cost so far and hasn't been processed yet...
		if cost < lowest_cost && !processed[node] {
			// ... set it as the new lowest-cost node.
			lowest_cost = cost
			lowest_cost_node = node
		}
	}
	return lowest_cost_node
}
