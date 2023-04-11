package historia_engine

import (
	"github.com/LuigiVanacore/AirWars2D/historia_engine/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	textureRect math2d.Rect
	texture     *ebiten.Image
	pivot       math2d.Vector2D
}

func NewSprite(texture *ebiten.Image) *Sprite {
	pivot := math2d.Vector2D{X: float64(texture.Bounds().Min.X),
		Y: float64(texture.Bounds().Min.Y)}
	textureRect := math2d.NewRect(float64(texture.Bounds().Min.X),
		float64(texture.Bounds().Min.Y),
		float64(texture.Bounds().Max.X),
		float64(texture.Bounds().Max.Y))

	return &Sprite{textureRect: *textureRect, texture: texture, pivot: pivot}
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

func (s *Sprite) GetPivot() math2d.Vector2D {
	return s.pivot
}

func (s *Sprite) SetPivot(x, y float64) {
	s.pivot = math2d.Vector2D{X: x, Y: y}
}

func (s *Sprite) SetPivotToCenter() {
	x, y := s.textureRect.GetCenter()
	s.pivot = math2d.Vector2D{X: x, Y: y}
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
