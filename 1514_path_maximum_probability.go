package main

import (
	"container/heap"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {

	graph := make([][]struct {
		node int
		prob float64
	}, n)
	for i, edge := range edges {
		a, b := edge[0], edge[1]
		prob := succProb[i]
		graph[a] = append(graph[a], struct {
			node int
			prob float64
		}{b, prob})
		graph[b] = append(graph[b], struct {
			node int
			prob float64
		}{a, prob})
	}

	type item struct {
		node int
		prob float64
	}
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, item{start, 1.0})

	probTo := make([]float64, n)
	probTo[start] = 1.0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(item)
		currNode, currProb := curr.node, curr.prob

		if currNode == end {
			return currProb
		}

		for _, neighbor := range graph[currNode] {
			nextNode, edgeProb := neighbor.node, neighbor.prob
			newProb := currProb * edgeProb
			if newProb > probTo[nextNode] {
				probTo[nextNode] = newProb
				heap.Push(pq, item{nextNode, newProb})
			}
		}
	}

	return 0.0
}

type priorityQueue []item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].prob > pq[j].prob
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
