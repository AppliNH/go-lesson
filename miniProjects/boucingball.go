package miniProjects

import (
	"fmt"
	"time"

	screen "github.com/inancgumus/screen"
)

type ball []int


func BouncingBall(x int, y int) {
	ClearScreen()
	var (
		py, px int
		vx, vy = 1, 1
		cell   rune
	)

	const (
		width     = 50
		height    = 10
		maxFrames = 1200
		cellEmpty = ' '
		cellBall  = '⚾'
		speed     = time.Second / 20
	)

	board := make([][]bool, width)

	for column := range board {
		board[column] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}
		board[px][py] = true

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {

				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Printf(string(buf))
		time.Sleep(speed)
	}
}
