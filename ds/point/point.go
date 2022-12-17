package point

import "fmt"

type Point struct {
	X, Y int
}
type Direction int

const (
	U Direction = 0
	R Direction = 1
	D Direction = 2
	L Direction = 3
)

func (p *Point) Id() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p *Point) Move(direction Direction) *Point {
	switch direction {
	case U:
		return NewPoint(p.Y-1, p.X)
	case R:
		return NewPoint(p.Y, p.X+1)
	case D:
		return NewPoint(p.Y+1, p.X)
	case L:
		return NewPoint(p.Y, p.X-1)
	default:
		fmt.Println("ho shit")
		return nil
	}
}

func (p *Point) MoveMutate(direction Direction) {
	switch direction {
	case U:
		p.Y--
	case R:
		p.X++
	case D:
		p.Y++
	case L:
		p.X--
	default:
		fmt.Println("ho shit")
	}
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}
