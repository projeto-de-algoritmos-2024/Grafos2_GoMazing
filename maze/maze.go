package maze

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Maze struct {
	TileSize   float64
	Rows       int
	Cols       int
	Grid       [][]int
	Directions [][2]int
	Win        *pixelgl.Window
}

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
		TileSize:   float64(tileSize),
		Rows:       rows,
		Cols:       cols,
		Grid:       grid,
		Directions: [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
		Win:        win,
	}
}

func (m *Maze) DrawMaze() {
	m.Win.Clear(colornames.Black)
	for y := 0; y < m.Rows; y++ {
		for x := 0; x < m.Cols; x++ {
			if m.Grid[y][x] == 1 {
				m.Win.SetColorMask(colornames.White)
				pixel.NewSprite(nil, pixel.R(0, 0, m.TileSize, m.TileSize)).Draw(m.Win, pixel.IM.Moved(pixel.V(float64(x)*m.TileSize+m.TileSize/2, float64(y)*m.TileSize+m.TileSize/2)))
			}
		}
	}
	m.Win.Update()
}

func (m *Maze) IsWithinBounds(x, y int) bool {
	return x >= 0 && x < m.Cols && y >= 0 && y < m.Rows
}

func (m *Maze) GenerateMaze(startX, startY int) {
	stack := [][2]int{{startX, startY}}
	m.Grid[startY][startX] = 1

	for len(stack) > 0 {
		x, y := stack[len(stack)-1][0], stack[len(stack)-1][1]
		directions := append([][2]int(nil), m.Directions...)
		rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

		foundPath := false
		for _, d := range directions {
			nx, ny := x+d[0]*2, y+d[1]*2
			if m.IsWithinBounds(nx, ny) && m.Grid[ny][nx] == 0 {
				m.Grid[y+d[1]][x+d[0]] = 1
				m.Grid[ny][nx] = 1
				stack = append(stack, [2]int{nx, ny})
				foundPath = true
				break
			}
		}

		if !foundPath {
			stack = stack[:len(stack)-1]
		}

		if len(stack)%500 == 0 {
			m.DrawMaze()
		}
	}
}

func (m *Maze) DrawCell(x, y int, color pixel.RGBA) {
	m.Win.SetColorMask(color)
	pixel.NewSprite(nil, pixel.R(0, 0, m.TileSize, m.TileSize)).Draw(m.Win, pixel.IM.Moved(pixel.V(float64(x)*m.TileSize+m.TileSize/2, float64(y)*m.TileSize+m.TileSize/2)))
	m.Win.Update()
}
