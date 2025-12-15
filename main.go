package main

import (
	"os"
	"time"
)

const w = 30
const h = 15

type Canvas [][]string

func ClearScreen() {
	os.Stdout.Write([]byte("\033[2J\033[H"))
}

func HideCursor() {
	os.Stdout.Write([]byte("\033[?25l"))
}

func Render() {

}

func Update() {

}

func main() {

	HideCursor()

	c := make(Canvas, w)
	for i := range len(c) {
		c[i] = make([]string, h)
	}

	for i := range len(c[0]) {
		for j := range len(c) {
			c[j][i] = "█"
		}
	}

	for {
		ClearScreen()
		for i := range len(c[0]) {
			for j := range len(c) {
				os.Stdout.Write([]byte(c[j][i]))
			}
			os.Stdout.Write([]byte("\n"))
		}
		time.Sleep(time.Millisecond * 12)
	}

}
