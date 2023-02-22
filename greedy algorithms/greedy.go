package main

import (
	"fmt"
)

func main() {
	statesNeeded := map[string]bool{
		"mt": true,
		"wa": true,
		"or": true,
		"id": true,
		"nv": true,
		"ut": true,
		"ca": true,
		"az": true,
	}

	stations := map[string]map[string]bool{
		"kone":   map[string]bool{"id": true, "nv": true, "ut": true},
		"ktwo":   map[string]bool{"wa": true, "id": true, "mt": true},
		"kthree": map[string]bool{"or": true, "nv": true, "ca": true},
		"kfour":  map[string]bool{"nv": true, "ut": true},
		"kfive":  map[string]bool{"ca": true, "az": true},
	}

	finalStations := make(map[string]bool)

	for len(statesNeeded) > 0 {
		bestStation := ""
		statesCovered := make(map[string]bool)

		for station, statesForStation := range stations {
			covered := make(map[string]bool)
			for state := range statesNeeded {
				if statesForStation[state] {
					covered[state] = true
				}
			}

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		if bestStation != "" {
			for state := range statesCovered {
				delete(statesNeeded, state)
			}

			finalStations[bestStation] = true
			delete(stations, bestStation)
		} else {
			break
		}
	}

	fmt.Println(finalStations)
}
