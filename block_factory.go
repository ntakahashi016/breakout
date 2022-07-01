package main

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type BlockFactory struct {
	game *Game
	height int
	width int
	hp int
}

func NewBlockFactory(g *Game, h,w,hp int) *BlockFactory {
	f := &BlockFactory{}
	f.game = g
	f.height = h
	f.width = w
	f.hp = hp
	return f
}

func (f *BlockFactory) NewObject(x,y int, color color.RGBA) *Block {
	image := ebiten.NewImage(f.width,f.height)
	image.Fill(color)
	o := Object{game: f.game, point: NewPoint(float64(x), float64(y)), height: f.height, width: f.width, image: image}
	return NewBlock(o, f.hp)
}
