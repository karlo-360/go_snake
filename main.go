package main

import (
	"fmt"
	"os"
	"time"
)

const w = 40
const h = 15

type Cell struct {
	Ch rune
}

type Canvas [][]Cell

func NewCanvas(width, height int) Canvas {
	c := make(Canvas, width)
	for i := range len(c) {
		c[i] = make([]Cell, height)
	}
	return c
}

func (c Canvas) FillCanvas() {
	for i := range len(c[0]) {
		for j := range len(c) {
			c[j][i] = Cell{Ch: '█'}
		}
	}
}

type Coord struct {
	X int
	Y int
}

type Snake struct {
	Head Coord
	Tail []Coord
	HeadCh rune
	TailCh rune
}

func NewSnake(x, y int) *Snake {
	return &Snake {
		Head: Coord{
			X: x,
			Y: y,
		},
		Tail: []Coord{
			{X: x-1, Y: y},
			{X: x-2, Y: y},
			{X: x-3, Y: y},
		},
		HeadCh: 'H',
		TailCh: 'T',

	}
}

func (s *Snake) UpdateSnake() {

	len := len(s.Tail)

	if len == 0 {
		return
	}

	for i := len - 1; i > 0; i-- {
		s.Tail[i] = s.Tail[i-1]
	}
	s.Tail[0] = s.Head
	s.Head.Y += 1
	s.Head.X += 1

}

type World struct {
	Canvas Canvas
	Snake *Snake
}

func NewWorld() *World {
	c := NewCanvas(w, h)
	c.FillCanvas()

	s := NewSnake(1, 1)

	return &World {
		Canvas: c,
		Snake: s,
	}
}

func (w *World) Render() {

	isCanvas := true

	ClearScreen()

	for i := range len(w.Canvas[0]) {
		for j := range len(w.Canvas) {


			for k := range len(w.Snake.Tail) {
				if w.Snake.Tail[k].X == j && w.Snake.Tail[k].Y  == i {
					fmt.Printf("%c", w.Snake.TailCh)
					isCanvas = false
				}
			}

			if w.Snake.Head.X == j && w.Snake.Head.Y == i {
				fmt.Printf("%c", w.Snake.HeadCh)
				isCanvas = false
			} 

			if isCanvas{
				fmt.Printf("%c", w.Canvas[j][i].Ch)
			}
			isCanvas = true
		}
		os.Stdout.Write([]byte("\n"))
	}
	time.Sleep(time.Millisecond * 200)
}

func (w *World) Update() {
	w.Snake.UpdateSnake()
}

func ClearScreen() {
	os.Stdout.Write([]byte("\033[2J\033[H"))
}

func HideCursor() {
	os.Stdout.Write([]byte("\033[?25l"))
}

func main() {

	HideCursor()

	world := NewWorld()

	for {
		world.Render()
		world.Update()
	}

}
