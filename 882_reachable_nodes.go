package main

import (
	"container/heap"
	"math"
)

type Edge struct {
	to, weight int
}

type Item struct {
	node, dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {

	adj := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		d, neigh := item.dist, item.node

		if d > dist[neigh] {
			continue
		}

		for _, edge := range adj[neigh] {
			v, count := edge.to, edge.weight
			newDist := d + count + 1
			if newDist < dist[v] && newDist <= maxMoves {
				dist[v] = newDist
				heap.Push(pq, &Item{v, newDist})
			}
		}
	}

	reachables := 0
	for i := 0; i < n; i++ {
		if dist[i] <= maxMoves {
			reachables++
		}
	}

	for _, e := range edges {
		u, v, count := e[0], e[1], e[2]
		canGoFromU := max(0, maxMoves-dist[u])
		canGoFromV := max(0, maxMoves-dist[v])
		reachables += min(count, canGoFromU+canGoFromV)
	}

	return reachables
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
