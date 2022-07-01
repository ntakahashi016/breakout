package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	game, _ := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Breakout")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
