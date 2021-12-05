package models

type Point struct {
	x int
	y int
}

func (point1 Point) hasSameXCoordinateAs(point2 Point) bool {
	return point1.x == point2.x
}

func (point1 Point) hasSameYCoordinateAs(point2 Point) bool {
	return point1.y == point2.y
}

func (point1 Point) isXCoordinateLesserOrEqualTo(point2 Point) bool {
	return point1.x <= point2.x
}

func (point1 Point) isYCoordinateLesserOrEqualTo(point2 Point) bool {
	return point1.y <= point2.y
}

func (point1 Point) isXCoordinateGreaterThan(point2 Point) bool {
	return point1.x > point2.x
}

func (point1 Point) isYCoordinateGreaterThan(point2 Point) bool {
	return point1.y > point2.y
}