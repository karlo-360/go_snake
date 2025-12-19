package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

const w = 40
const h = 15

type Movement int
const (
	Left Movement = iota
	Right
	Up
	Down
	No
)

type Cell struct {
	Ch rune
}

type Canvas [][]Cell

func NewCanvas(width, height int) *Canvas {
	c := make(Canvas, width)
	for i := range len(c) {
		c[i] = make([]Cell, height)
	}
	return &c
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

type Input byte

type Snake struct {
	Head Coord
	Tail []Coord
	HeadCh rune
	TailCh rune
	Direction Movement
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
		Direction: Right,
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

}

type World struct {
	Canvas *Canvas
	Snake *Snake
	Input Input
	IsRunning bool
}

func NewWorld() *World {
	c := NewCanvas(w, h)
	c.FillCanvas()

	s := NewSnake(5, 1)

	return &World {
		Canvas: c,
		Snake: s,
		IsRunning: true,
	}
}

func (w *World) Render() {

	isCanvas := true

	ClearScreen()

	c := *w.Canvas
	for i := range len(c[0]) {
		for j := range len(c) {


			for k := range len(w.Snake.Tail) {
				if w.Snake.Tail[k].X == j && w.Snake.Tail[k].Y  == i {
					fmt.Printf("%c", w.Snake.TailCh)
					isCanvas = false
					break
				}
			}

			if w.Snake.Head.X == j && w.Snake.Head.Y == i {
				fmt.Printf("%c", w.Snake.HeadCh)
				isCanvas = false
			} 

			if isCanvas{
				fmt.Printf("%c", c[j][i].Ch)
			}
			isCanvas = true
		}
		os.Stdout.Write([]byte("\r\n"))
	}
	time.Sleep(time.Millisecond * 200)
}

func (w *World) Update() {


	w.Snake.UpdateSnake()
	w.Snake.UpdateDirection()
}

func (s *Snake) UpdateDirection() {

	if s.Direction == Right {
	  s.Head.X += 1
	}
	
	if s.Direction == Left {
	  s.Head.X -= 1
	}

	if s.Direction == Up {
	  s.Head.Y -= 1
	}

	if s.Direction == Down {
	  s.Head.Y += 1
	}
}

func  RenderInfo() {
}

func ReadInput(out chan <- byte) {

	buf := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Printf("Failed to reading a key: %v\n", err)
			return 
		}
		out <- buf[0]
	}
}

func TickLoop(out chan <- struct{}) {
	ticker := time.NewTicker(2 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		out <- struct{}{}
	}
}

func (s *Snake) ReadMovement(input <- chan byte, tick <- chan struct{}) {
	select {
	case b := <- input:
		switch b {
		case 'i':
			s.Direction = Up
		case 'j':
			s.Direction = Left
		case 'k':
			s.Direction = Down
		case 'l':
			s.Direction = Right
		case 'q':
			os.Exit(0)
		}
	case <- tick:
	}
}

func ClearScreen() {
	os.Stdout.Write([]byte("\033[2J\033[H"))
}

func HideCursor() {
	os.Stdout.Write([]byte("\033[?25l"))
}

func RawMode() func() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("Failed to enable raw mode: %v\n", err)
	}
	return func() {
		defer term.Restore(int(os.Stdin.Fd()), oldState)
		if err != nil {
			fmt.Printf("Failed to restore normal mode: %v\n", err)
		}
	}
}

func main() {
	restore := RawMode()
	defer restore()

	HideCursor()
	i := make(chan byte)
	tick := make(chan struct{})

	go ReadInput(i)
	go TickLoop(tick)

	world := NewWorld()

	for world.IsRunning {

		println("ho")

		world.Snake.ReadMovement(i, tick)
		world.Render()
		world.Update()
	}

}
