package main

import (
	"fmt"
	"math"
)

func main() {
	graph := make(map[string]map[string]float64)
	graph["start"] = map[string]float64{"a": 6, "b": 2}
	graph["a"] = map[string]float64{"fin": 1}
	graph["b"] = map[string]float64{"a": 3, "fin": 5}
	graph["fin"] = map[string]float64{}

	infinity := math.Inf(1)

	costs := map[string]float64{"a": 6, "b": 2, "fin": infinity}

	parents := map[string]string{"a": "start", "b": "start", "fin": ""}

	processed := make(map[string]bool)

	findLowestCostNode := func(costs map[string]float64) string {
		lowestCost := math.Inf(1)
		lowestCostNode := ""
		for node, cost := range costs {
			if cost < lowestCost && !processed[node] {
				lowestCost = cost
				lowestCostNode = node
			}
		}
		return lowestCostNode
	}

	for node := findLowestCostNode(costs); node != ""; node = findLowestCostNode(costs) {
		cost := costs[node]
		neighbors := graph[node]
		for n, _ := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed[node] = true
	}

	fmt.Println("Cost from the start to each node:")
	fmt.Println(costs)
}
