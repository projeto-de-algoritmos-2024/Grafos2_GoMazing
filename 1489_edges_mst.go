package main

import (
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	type Edge struct {
		u, v, weight, index int
	}

	edgeList := make([]Edge, len(edges))
	for i, e := range edges {
		edgeList[i] = Edge{e[0], e[1], e[2], i}
	}

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i].weight < edgeList[j].weight
	})

	parent := make([]int, n)
	rank := make([]int, n)

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			if rank[rootX] > rank[rootY] {
				parent[rootY] = rootX
			} else if rank[rootX] < rank[rootY] {
				parent[rootX] = rootY
			} else {
				parent[rootY] = rootX
				rank[rootX]++
			}
			return true
		}
		return false
	}

	calculateMST := func(exclude, include int) int {
		for i := range parent {
			parent[i] = i
			rank[i] = 0
		}
		totalWeight := 0
		edgesUsed := 0

		if include != -1 {
			e := edgeList[include]
			if union(e.u, e.v) {
				totalWeight += e.weight
				edgesUsed++
			}
		}

		for i, e := range edgeList {
			if i == exclude {
				continue
			}
			if union(e.u, e.v) {
				totalWeight += e.weight
				edgesUsed++
				if edgesUsed == n-1 {
					break
				}
			}
		}

		if edgesUsed == n-1 {
			return totalWeight
		}
		return 1<<31 - 1
	}

	originalMSTWeight := calculateMST(-1, -1)

	criticalEdges := []int{}
	pseudoCriticalEdges := []int{}

	for i := range edgeList {

		if calculateMST(i, -1) > originalMSTWeight {
			criticalEdges = append(criticalEdges, edgeList[i].index)
		} else if calculateMST(-1, i) == originalMSTWeight {

			pseudoCriticalEdges = append(pseudoCriticalEdges, edgeList[i].index)
		}
	}

	return [][]int{criticalEdges, pseudoCriticalEdges}
}
