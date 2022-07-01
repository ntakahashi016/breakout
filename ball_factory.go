package main

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type BallFactory struct {
	game *Game
	image *ebiten.Image
	height int
	width int
}

func NewBallFactory(g *Game, h,w int) *BallFactory {
	f := &BallFactory{}
	f.game = g
	f.image = ebiten.NewImage(w,h)
	f.image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	f.height = h
	f.width = w
	return f
}

func (f *BallFactory) NewObject(x,y int) *Ball {
	o := Object{game: f.game, point: NewPoint(float64(x), float64(y)), height: f.height, width: f.width, image: f.image}
	return NewBall(o)
}
