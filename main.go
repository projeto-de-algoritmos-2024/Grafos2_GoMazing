package main

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/projeto-de-algoritmos-2024/Grafos2_GoMazing/algorithms"
	"github.com/projeto-de-algoritmos-2024/Grafos2_GoMazing/maze"
)

func run() {
	width, height := 500, 500
	tileSize := 5

	m := maze.NewMaze(width, height, tileSize)
	m.GenerateMaze(0, 0)
	m.DrawMaze()

	start := [2]int{0, 0}
	end := [2]int{m.Cols - 2, m.Rows - 2}

	algorithms.Dijkstra(m, start, end)

	time.Sleep(10 * time.Second)
}

func main() {
	pixelgl.Run(run)
}
