package algorithms

import (
	"fmt"
	"time"

	"github.com/projeto-de-algoritmos-2024/Grafos2_GoMazing/main"
)

func (l *main.Maze) DFS(start, end [2]int) {
	stack := [][2]int{start}
	visited := make(map[[2]int]bool)
	ticker := time.NewTicker(time.Millisecond * 2)
	defer ticker.Stop()

	for len(stack) > 0 {
		select {
		case <-ticker.C:
			if l.win.Closed() {
				return
			}

			x, y := stack[len(stack)-1][0], stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			if visited[[2]int{x, y}] {
				continue
			}
			visited[[2]int{x, y}] = true

			l.drawPath(x, y)

			if [2]int{x, y} == end {
				fmt.Println("Ponto de destino encontrado!")
				return
			}

			if l.grid[y][x] == 1 {
				for _, d := range l.directions {
					nx, ny := x+d[0], y+d[1]
					if l.isWithinBounds(nx, ny) && l.grid[ny][nx] == 1 && !visited[[2]int{nx, ny}] {
						stack = append(stack, [2]int{nx, ny})
					}
				}
			}
		}
	}
}
