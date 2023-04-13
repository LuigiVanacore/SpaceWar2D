package historia_engine

import "github.com/hajimehoshi/ebiten/v2"

type EntitySprite struct {
	Entity2D
	sprite *Sprite
}

func NewEntitySprite(sprite *Sprite) *EntitySprite {
	return &EntitySprite{sprite: sprite}
}

func (e *EntitySprite) Update() {

}

func (e *EntitySprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM = e.geoM
	e.sprite.Draw(target, op)
}

func (e *EntitySprite) SetPivotToCenter() {
	rect := e.sprite.GetTextureRect()
	x, y := rect.GetCenter()
	e.SetPivot(x, y)
}
