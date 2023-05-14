package point

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetCoords() [2]float64 {
	return [2]float64{p.x, p.y}
}
