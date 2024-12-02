package main

import (
	"container/heap"
)

func reachableNodes(edges [][]int, maxMoves int, n int) int {

	graph := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, cnt := edge[0], edge[1], edge[2]
		if graph[u] == nil {
			graph[u] = make(map[int]int)
		}
		if graph[v] == nil {
			graph[v] = make(map[int]int)
		}
		graph[u][v] = cnt
		graph[v][u] = cnt
	}

	type Item struct {
		node, moves int
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, 0})

	minMoves := make([]int, n)
	for i := range minMoves {
		minMoves[i] = maxMoves + 1
	}
	minMoves[0] = 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node, moves := item.node, item.moves

		if moves > minMoves[node] {
			continue
		}

		for neighbor, cnt := range graph[node] {
			newMoves := moves + cnt + 1
			if newMoves < minMoves[neighbor] {
				minMoves[neighbor] = newMoves
				heap.Push(pq, &Item{neighbor, newMoves})
			}
		}
	}

	reachable := 0
	for _, moves := range minMoves {
		if moves <= maxMoves {
			reachable++
		}
	}

	for _, edge := range edges {
		u, v, cnt := edge[0], edge[1], edge[2]
		reachable += min(cnt, max(0, maxMoves-minMoves[u])) + min(cnt, max(0, maxMoves-minMoves[v]))
	}

	return reachable
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].moves < pq[j].moves
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
