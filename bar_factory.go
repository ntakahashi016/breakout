package main

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type BarFactory struct {
	game *Game
	image *ebiten.Image
	height int
	width int
}

func NewBarFactory(g *Game, h,w int) *BarFactory {
	f := &BarFactory{}
	f.game = g
	f.image = ebiten.NewImage(w,h)
	f.image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	f.height = h
	f.width = w
	return f
}

func (f *BarFactory) NewObject(x,y int) *Bar {
	o := Object{game: f.game, point: NewPoint(float64(x), float64(y)), height: f.height, width: f.width, image: f.image}
	i := NewInput()
	return NewBar(o, i)
}
