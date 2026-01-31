package sprint


type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point) Point {
	d1 := p1.X*p1.X + p1.Y*p1.Y
	d2 := p2.X*p2.X + p2.Y*p2.Y

	if d1 >= d2 {
		return p1
	}
	return p2
}
