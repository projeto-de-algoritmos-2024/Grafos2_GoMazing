package main

import (
	"github.com/faiface/pixel/pixelgl"
)

type Maze struct {
	tileSize   float64
	rows       int
	cols       int
	grid       [][]int
	directions [][2]int
	win        *pixelgl.Window
}
