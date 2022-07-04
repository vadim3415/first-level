package pointer
// пакет от 24 задания

type point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func (p *point) Get() (x, y float64){
	return p.x, p.y
}