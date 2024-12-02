package main

import (
	"container/heap"
)

type Item struct {
	node, moves int
}

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

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, maxMoves})

	visited := make(map[int]int)

	reachable := 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node, moves := item.node, item.moves

		if remaining, seen := visited[node]; seen && remaining >= moves {
			continue
		}
		visited[node] = moves
		reachable++

		if moves > 0 {
			for neighbor, cnt := range graph[node] {
				usedMoves := min(moves, cnt)
				graph[node][neighbor] -= usedMoves
				graph[neighbor][node] -= usedMoves
				reachable += usedMoves

				if moves > cnt {
					if remaining, seen := visited[neighbor]; !seen || remaining < moves-cnt-1 {
						heap.Push(pq, &Item{neighbor, moves - cnt - 1})
					}
				}
			}
		}
	}

	for _, edge := range edges {
		u, v, cnt := edge[0], edge[1], edge[2]
		reachable += min(cnt, visited[u]+visited[v])
	}

	return reachable
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].moves > pq[j].moves
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
