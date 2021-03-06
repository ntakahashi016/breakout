package main

type HitDirection int
const (
	None HitDirection = iota
	Vertical
	Horizontal
)

type Area struct {
	p1 *Point
	p2 *Point
}

func NewArea(p1,p2 *Point) *Area {
	a := &Area{}
	var x1,y1,x2,y2 float64
	if (p2.x - p1.x) >= 0 {
		x1 = p1.x
		x2 = p2.x
	} else {
		x1 = p2.x
		x2 = p1.x
	}
	if (p2.y - p1.y) >= 0 {
		y1 = p1.y
		y2 = p2.y
	} else {
		y1 = p2.y
		y2 = p1.y
	}
	a.p1 = NewPoint(x1,y1)
	a.p2 = NewPoint(x2,y2)
	return a
}

func (a *Area) isHit(other *Area) HitDirection {
	if a.p1.x <= other.p2.x && other.p1.x <= a.p2.x && a.p1.y <= other.p2.y && other.p1.y <= a.p2.y {
		if a.p1.x <= other.p1.x && a.p1.y <= other.p1.y {
			if a.p2.x - other.p1.x <= a.p2.y - other.p1.y {
				return Horizontal
			} else {
				return Vertical
			}
		} else if a.p1.x <= other.p1.x && other.p1.y <= a.p1.y {
			if a.p2.x - other.p1.x <= other.p2.y - a.p1.y {
				return Horizontal
			} else {
				return Vertical
			}
		} else if other.p1.x <= a.p1.x && a.p1.y <= other.p1.y {
			if other.p2.x - a.p1.x <= a.p2.y - other.p1.y {
				return Horizontal
			} else {
				return Vertical
			}
		} else {
			if other.p2.x - a.p1.x <= other.p2.y - a.p1.y {
				return Horizontal
			} else {
				return Vertical
			}
		}
	}
	return None
}

func (a *Area) offset(v Vector) *Area {
	return NewArea(a.p1.offset(v), a.p2.offset(v))
}
