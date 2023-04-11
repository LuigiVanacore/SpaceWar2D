package main

import (
	"github.com/LuigiVanacore/AirWars2D"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(airwar2d.WindowWidth, airwar2d.WindowHeight)
	ebiten.SetWindowTitle("AirWars2D")
	game := &airwar2d.Game{}
	game.Init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
