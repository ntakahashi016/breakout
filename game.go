package main

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	objects []interface{}
	ball *Ball
	clear bool
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.clear = false
	g.objects = []interface{}{}
	// left wall
	i := ebiten.NewImage(10, screenHeight)
	i.Fill(color.RGBA{0x55, 0x55, 0x55, 0xff})
	o := Object{game: g, point: NewPoint(0.0, 0.0), height: screenHeight, width: 10, image: i}
	g.objects = append(g.objects, NewWall(o))
	// upper wall
	i = ebiten.NewImage(screenWidth, 10)
	i.Fill(color.RGBA{0x55, 0x55, 0x55, 0xff})
	o = Object{game: g, point: NewPoint(0.0, 0.0), height: 10, width: screenWidth, image: i}
	g.objects = append(g.objects, NewWall(o))
	// right wall
	i = ebiten.NewImage(130, screenHeight)
	i.Fill(color.RGBA{0x55, 0x55, 0x55, 0xff})
	o = Object{game: g, point: NewPoint(screenWidth-130.0, 0.0), height: screenHeight, width: 20, image: i}
	g.objects = append(g.objects, NewWall(o))
	// blocks
	w := 50
	h := 10
	hp := 1
	blockFactory := NewBlockFactory(g,h,w,hp)
	for y := 0; y < screenHeight/2; y+=h  {
		for x := 0; x < screenWidth-140; x+=w {
			red := (uint8)(0xff)
			green := (uint8)(0xff - (x/w)*10)
			blue := (uint8)(0xff - (y/h)*10)
			alpha := (uint8)(0xff)
			color := color.RGBA{red,green,blue,alpha}
			g.objects = append(g.objects, blockFactory.NewObject(x+10,y+10,color))
		}
	}
	// bar
	barFactory := NewBarFactory(g,1,100)
	g.objects = append(g.objects, barFactory.NewObject((screenWidth - 140 - 100)/2, screenHeight - 50))
	// ball
	ballFactory := NewBallFactory(g,10,10)
	g.ball = ballFactory.NewObject((screenWidth - 140 - 10)/2, screenHeight - 100)
	g.objects = append(g.objects, g.ball)
	return g, nil
}

func (g *Game) deleteObject(o interface{}) {
	newObjects := []interface{}{}
	for _,v := range g.objects{
		if v != o {
			newObjects = append(newObjects, v)
		}
	}
	g.objects = newObjects
}

func (g *Game) outOfScreen(a *Area) bool {
	return (a.p2.x < 0 || screenWidth <= a.p1.x || a.p2.y < 0 || screenHeight <= a.p1.y)
}

func (g *Game) insideOfScreen(a *Area) bool {
	return (a.p1.x >= 0 && screenWidth > a.p2.x && a.p1.y >= 0 && screenHeight > a.p2.y)
}

func (g *Game) repointOnScreen(a *Area) *Point {
	p := &Point{x: a.p1.x, y: a.p1.y}
	if a.p1.x < 0 {
		p.x = 0
	} else if screenWidth <= a.p2.x {
		p.x = screenWidth - (a.p2.x - a.p1.x)
	}
	if a.p1.y < 0 {
		p.y = 0
	} else if screenHeight <= a.p2.y {
		p.y = screenHeight - (a.p2.y - a.p1.y)
	}
	return p
}

func (g *Game) Update() error {
	if g.ball == nil { return nil }
	for _,v := range g.objects {
		c := v.(Common)
		c.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _,v := range g.objects {
		c := v.(Common)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.X()), float64(c.Y()))
		screen.DrawImage(c.Image(), op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
