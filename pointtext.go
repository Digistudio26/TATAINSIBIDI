package main

type Point struct {
	X    int
	Y    int
	Text string
}

func PointText(p Point) Point {
	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: fmt.Sprintf("Text at (%d, %d)", p.X, p.Y),
	}
}
