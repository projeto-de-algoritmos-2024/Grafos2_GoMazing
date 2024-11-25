package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/projeto-de-algoritmos-2024/Grafos2_GoMazing/algorithms"
	"golang.org/x/image/colornames"
)

func NewMaze(width, height, tileSize int) *Maze {
	rows := height / tileSize
	cols := width / tileSize
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	cfg := pixelgl.WindowConfig{
		Title:  "Grafos 2 Maze",
		Bounds: pixel.R(0, 0, float64(width), float64(height)),
		VSync:  true,
	}
	win, _ := pixelgl.NewWindow(cfg)
	return &Maze{
		tileSize:   float64(tileSize),
		rows:       rows,
		cols:       cols,
		grid:       grid,
		directions: [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
		win:        win,
	}
}

func (l *Maze) drawMaze() {
	l.win.Clear(colornames.Black)
	for y := 0; y < l.rows; y++ {
		for x := 0; x < l.cols; x++ {
			if l.grid[y][x] == 1 {
				l.win.SetColorMask(colornames.White)
				pixel.NewSprite(nil, pixel.R(0, 0, l.tileSize, l.tileSize)).Draw(l.win, pixel.IM.Moved(pixel.V(float64(x)*l.tileSize+l.tileSize/2, float64(y)*l.tileSize+l.tileSize/2)))
			}
		}
	}
	l.win.Update()
}

func (l *Maze) isWithinBounds(x, y int) bool {
	return x >= 0 && x < l.cols && y >= 0 && y < l.rows
}

func (l *Maze) generateMaze(startX, startY int) {
	stack := [][2]int{{startX, startY}}
	l.grid[startY][startX] = 1

	for len(stack) > 0 {
		x, y := stack[len(stack)-1][0], stack[len(stack)-1][1]
		directions := append([][2]int(nil), l.directions...)
		rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

		foundPath := false
		for _, d := range directions {
			nx, ny := x+d[0]*2, y+d[1]*2
			if l.isWithinBounds(nx, ny) && l.grid[ny][nx] == 0 {
				l.grid[y+d[1]][x+d[0]] = 1
				l.grid[ny][nx] = 1
				stack = append(stack, [2]int{nx, ny})
				foundPath = true
				break
			}
		}

		if !foundPath {
			stack = stack[:len(stack)-1]
		}

		if len(stack)%500 == 0 {
			l.drawMaze()
		}
	}
}

func (l *Maze) drawPath(x, y int) {
	l.win.SetColorMask(colornames.Blue)
	pixel.NewSprite(nil, pixel.R(0, 0, l.tileSize, l.tileSize)).Draw(l.win, pixel.IM.Moved(pixel.V(float64(x)*l.tileSize+l.tileSize/2, float64(y)*l.tileSize+l.tileSize/2)))
	l.win.Update()
}

func (l *Maze) drawCell(x, y int, color pixel.RGBA) {
	l.win.SetColorMask(color)
	pixel.NewSprite(nil, pixel.R(0, 0, l.tileSize, l.tileSize)).Draw(l.win, pixel.IM.Moved(pixel.V(float64(x)*l.tileSize+l.tileSize/2, float64(y)*l.tileSize+l.tileSize/2)))
	l.win.Update()
}

func run() {
	width, height := 500, 500
	tileSize := 5

	maze := NewMaze(width, height, tileSize)
	maze.generateMaze(0, 0)
	maze.drawMaze()
	start := [2]int{0, 0}
	end := [2]int{maze.cols - 2, maze.rows - 2}

	algorithm := "Dijkstra" // Change to "DFS" to use DFS algorithm
	if algorithm == "Dijkstra" {
		algorithms.Dijkstra(maze, start, end)
	} else if algorithm == "DFS" {
		algorithms.DFS(maze, start, end)
	}

	time.Sleep(10 * time.Second)
}

func main() {
	pixelgl.Run(run)
}
