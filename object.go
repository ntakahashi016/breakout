package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	game *Game
	point *Point
	height int
	width int
	image *ebiten.Image
}

type Common interface {
	Update()
	Draw(img *ebiten.Image) error
	X() int
	Y() int
	Area() *Area
	Image() *ebiten.Image
	Center() *Point
	Hit(int)
}
