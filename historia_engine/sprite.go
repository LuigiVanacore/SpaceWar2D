package historia_engine

import (
	"github.com/LuigiVanacore/AirWars2D/historia_engine/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	textureRect math2d.Rect
	texture     *ebiten.Image
}

func NewSprite(texture *ebiten.Image) *Sprite {

	textureRect := math2d.NewRect(float64(texture.Bounds().Min.X),
		float64(texture.Bounds().Min.Y),
		float64(texture.Bounds().Max.X),
		float64(texture.Bounds().Max.Y))

	return &Sprite{textureRect: *textureRect, texture: texture}
}

func (s *Sprite) GetTextureRect() math2d.Rect {
	return s.textureRect
}

func (s *Sprite) SetTextureRect(width, height float64) {
	s.textureRect = math2d.Rect{Width: width, Height: height}
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}
func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

//
//func (s* Sprite) GetLocalBounds() Math.Rect {
//	return s.textureRect
//}
//
//func (s* Sprite) GetGlobalBounds() Math.Rect {
//	transform := s.GetTransform()
//	return transform.(s.GetLocalBounds())
//}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {

	if s.texture != nil {
		//s.transform.Concat(op.GeoM)
		target.DrawImage(s.texture, op)
	}
}
