package main





type Point struct {
	X    float64
	Y    float64
	Text string
}

func PointText(p Point) Point {
	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: fmt.Sprint("Text at (", p.X, ", ", p.Y, ")"),
	}
}
