package algorithms

import (
	"container/heap"
	"time"

	"github.com/faiface/pixel"
	"github.com/projeto-de-algoritmos-2024/Grafos2_GoMazing/main"
	"golang.org/x/image/colornames"
)

func (l *main.Maze) dijkstra(start, end [2]int) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{priority: 0, point: start})
	distances := map[[2]int]int{start: 0}
	previous := map[[2]int][2]int{}
	visited := make(map[[2]int]bool)

	for pq.Len() > 0 {
		if l.win.Closed() {
			return
		}

		currentItem := heap.Pop(pq).(*PriorityQueueItem)
		current := currentItem.point
		if visited[current] {
			continue
		}
		visited[current] = true

		x, y := current[0], current[1]
		l.drawCell(x, y, pixel.ToRGBA(colornames.Green)) // Convert color.RGBA to pixel.RGBA
		time.Sleep(10 * time.Millisecond)

		if current == end {
			path := [][2]int{}
			for current != start {
				path = append(path, current)
				current = previous[current]
			}
			path = append(path, start)
			for i := len(path) - 1; i >= 0; i-- {
				l.drawCell(path[i][0], path[i][1], pixel.ToRGBA(colornames.Blue)) // Convert color.RGBA to pixel.RGBA
				time.Sleep(5 * time.Millisecond)
			}
			return
		}

		for _, d := range l.directions {
			nx, ny := x+d[0], y+d[1]
			if l.isWithinBounds(nx, ny) && l.grid[ny][nx] == 1 {
				neighbor := [2]int{nx, ny}
				newDistance := distances[current] + 1
				if _, ok := distances[neighbor]; !ok || newDistance < distances[neighbor] {
					distances[neighbor] = newDistance
					previous[neighbor] = current
					heap.Push(pq, &PriorityQueueItem{priority: newDistance, point: neighbor})
				}
			}
		}
	}
}
