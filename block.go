package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Block struct {
	Common
	Object
	hp int
}

func NewBlock(o Object, hp int) *Block {
	b := &Block{
		Object: o,
	}
	b.hp = hp
	return b
}

func (b *Block) Update() {}
func (b *Block) X() int { return int(b.point.X()) }
func (b *Block) Y() int { return int(b.point.Y()) }
func (b *Block) Area() *Area {
	return NewArea(NewPoint(b.point.X(), b.point.Y()), NewPoint(b.point.X()+float64(b.width), b.point.Y()+float64(b.height)) )
}
func (b *Block) Image() *ebiten.Image { return b.image }
func (b *Block) Center() *Point {
	a := b.Area()
	x := (a.p2.x - a.p1.x) / 2 + a.p1.x
	y := (a.p2.y - a.p1.y) / 2 + a.p1.y
	return NewPoint(x,y)
}
func (b *Block) Hit(damage int) {
	b.hp -= damage
	if b.hp <= 0 {
		b.destroy()
	}
}
func (b *Block) destroy() {
	b.game.deleteObject(b)
}
