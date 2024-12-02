package main

import (
	"container/heap"
)

type Item struct {
	x, y, cost int
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	directions := []struct {
		dx, dy int
	}{
		{0, 1},  // right
		{0, -1}, // left
		{1, 0},  // down
		{-1, 0}, // up
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{0, 0, 0})

	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = 1<<31 - 1
		}
	}
	costs[0][0] = 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		x, y, cost := item.x, item.y, item.cost

		if x == m-1 && y == n-1 {
			return cost
		}

		for i, dir := range directions {
			nx, ny := x+dir.dx, y+dir.dy
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				newCost := cost
				if grid[x][y] != i+1 {
					newCost++
				}
				if newCost < costs[nx][ny] {
					costs[nx][ny] = newCost
					heap.Push(pq, &Item{nx, ny, newCost})
				}
			}
		}
	}

	return costs[m-1][n-1]
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
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
