package historia_engine

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}
