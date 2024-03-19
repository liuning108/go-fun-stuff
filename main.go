package main

import (
	"fmt"
)

type Position struct {
	x, y int
}

func (p *Position) Move(x int, y int) string {
	p.x = x
	p.y = y
	return fmt.Sprintf("The position is moverd by  : %d, %d", p.x, p.y)
}

type Player struct {
	Position
}

type Color int

func (c Color) String() string {
	return [...]string{"BLUE", "BLACK", "YELLOW", "PINK"}[c]
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	p := Player{}
	s := p.Move(10, 20)
	fmt.Println(s)
	fmt.Println("the color is : ", ColorBlue)
}
