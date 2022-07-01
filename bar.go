package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bar struct {
	Common
	Object
	input *Input
	moved bool
	vector Vector
}

func NewBar(o Object, input *Input) *Bar {
	b := &Bar{
		Object: o,
	}
	b.input = input
	b.moved = false
	b.vector = NewVector(0.0, 0.0)
	return b
}
func (b *Bar) command(cmd Command) {
	switch cmd {
	case KeySpace:
	}
}
func (b *Bar) move(v Vector) bool {
	moved := false
	b.vector = v
	na := b.Area().offset(v)
	if b.game.insideOfScreen(na) {
		if v.X() != 0.0 {
			b.point = na.p1
			moved = true
		} else {
			moved = false
		}
	} else {
		rp := b.game.repointOnScreen(na)
		b.point = rp
		moved = false
	}
	return moved
}

func (b *Bar) Update() {
	commands := b.input.getCommands()
	for len(commands) > 0 {
		command := commands[0]
		commands = commands[1:]
		b.command(command)
	}
	vector := b.input.getVector()
	hitAreas := []*Area{}
	nv := vector.Normalize()
	from := b.point
	destination := b.point.offset(vector)
	for !from.equal(destination){
		np1 := b.point.offset(nv)
		np2 := np1.offset(NewVector(float64(b.width), float64(b.height)))
		na  := NewArea(np1, np2)
		hitAreas = append(hitAreas, na)
		from = np1
	}
	for _, o := range b.game.objects {
		switch o := o.(type) {
		case *Wall:
			for _,v := range hitAreas {
				switch v.isHit(o.Area()) {
				case Horizontal:
					vector = NewVector(0.0, 0.0)
				default:
				}
			}
		case Common:
		}
	}
	b.moved = b.move(vector)

}
func (b *Bar) X() int { return int(b.point.X()) }
func (b *Bar) Y() int { return int(b.point.Y()) }
func (b *Bar) Area() *Area {
	return NewArea(NewPoint(b.point.X(), b.point.Y()), NewPoint(b.point.X()+float64(b.width), b.point.Y()+float64(b.height)) )
}
func (b *Bar) Image() *ebiten.Image { return b.image }
func (b *Bar) Center() *Point {
	a := b.Area()
	x := (a.p2.x - a.p1.x) / 2 + a.p1.x
	y := (a.p2.y - a.p1.y) / 2 + a.p1.y
	return NewPoint(x,y)
}
func (b *Bar) Hit(damage int) {
}
func (b *Bar) isMoved() bool { return b.moved }
