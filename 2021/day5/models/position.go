package models

type Position struct {
	startPoint Point
	endPoint   Point
}

func NewPosition(x1 int, y1 int, x2 int, y2 int) Position {
	return Position{
		startPoint: Point{x: x1, y: y1},
		endPoint:   Point{x: x2, y: y2},
	}
}