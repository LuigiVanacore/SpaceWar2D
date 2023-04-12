package historia_engine

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Update()
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}
