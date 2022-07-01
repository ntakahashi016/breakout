package main

import  (
	"github.com/hajimehoshi/ebiten/v2"
)

type Wall struct {
	Common
	Object
}

func NewWall(o Object) *Wall {
	w := &Wall{
		Object: o,
	}
	return w
}
func (w *Wall) Update() {}
func (w *Wall) X() int { return int(w.point.X()) }
func (w *Wall) Y() int { return int(w.point.Y()) }
func (w *Wall) Area() *Area {
	return NewArea(NewPoint(w.point.X(), w.point.Y()), NewPoint(w.point.X()+float64(w.width), w.point.Y()+float64(w.height)) )
}
func (w *Wall) Image() *ebiten.Image { return w.image }
func (w *Wall) Center() *Point {
	a := w.Area()
	x := (a.p2.x - a.p1.x) / 2 + a.p1.x
	y := (a.p2.y - a.p1.y) / 2 + a.p1.y
	return NewPoint(x,y)
}
func (w *Wall) Hit(damage int) {
}
