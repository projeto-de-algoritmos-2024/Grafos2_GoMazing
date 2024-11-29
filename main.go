package main

import (
	"net/http"
	"time"

	"github.com/faiface/mainthread"
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
	mainthread.Run(func() {
		http.HandleFunc("/generate-maze", func(w http.ResponseWriter, r *http.Request) {
			mainthread.Call(run)
			w.Write([]byte("Maze generated"))
		})

		http.ListenAndServe(":8080", nil)
	})
}
