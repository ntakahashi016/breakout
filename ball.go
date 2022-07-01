package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	Common
	Object
	vector Vector
}

const MaxSpeed = 2.0

func NewBall(o Object) *Ball {
	b := &Ball{
		Object: o,
	}
	b.vector = NewVector(0.0,1.0)
	return b
}

func (b *Ball) Update() {
	if b.game.outOfScreen(b.Area()) {
		b.game.deleteObject(b)
		return
	}
	hitAreas := []*Area{}
	vector := b.vector.Normalize()
	destination := b.point.offset(b.vector)
	for !b.point.equal(destination) {
		np1 := b.point.offset(vector)
		np2 := np1.offset(NewVector(float64(b.width), float64(b.height)))
		na  := NewArea(np1, np2)
		hitAreas = append(hitAreas, na)
		b.point = np1
	}
	for _, o := range b.game.objects {
		switch o := o.(type) {
		case *Ball:
			break
		case *Bar:
			for _,v := range hitAreas {
				switch v.isHit(o.Area()) {
				case Vertical:
					o.Hit(1)
					if o.isMoved() {
						b.vector.SetX(b.vector.X() + o.vector.X())
						if b.vector.X() > MaxSpeed {
							b.vector.SetX(MaxSpeed)
						} else if -MaxSpeed > b.vector.X() {
							b.vector.SetX(-MaxSpeed)
						}
					}
					b.vector.SetY(-b.vector.Y())
				case Horizontal:
					o.Hit(1)
					b.vector.SetX(-b.vector.X())
				}
			}
		case Common:
			for _,v := range hitAreas {
				switch v.isHit(o.Area()) {
				case Vertical:
					o.Hit(1)
					b.vector.SetY(-b.vector.Y())
				case Horizontal:
					o.Hit(1)
					b.vector.SetX(-b.vector.X())
				}
			}
		}
	}
}
func (b *Ball) X() int { return int(b.point.X()) }
func (b *Ball) Y() int { return int(b.point.Y()) }
func (b *Ball) Area() *Area {
	return NewArea(NewPoint(b.point.X(), b.point.Y()), NewPoint(b.point.X()+float64(b.width), b.point.Y()+float64(b.height)) )
}
func (b *Ball) Image() *ebiten.Image { return b.image }
func (b *Ball) Center() *Point {
	a := b.Area()
	x := (a.p2.x - a.p1.x) / 2 + a.p1.x
	y := (a.p2.y - a.p1.y) / 2 + a.p1.y
	return NewPoint(x,y)
}
func (b *Ball) Hit(damage int) {
}
func (b *Ball) destroy() {
	b.game.deleteObject(b)
}
