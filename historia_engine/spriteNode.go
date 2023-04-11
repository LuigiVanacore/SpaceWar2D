package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteNode struct {
	SceneNode
	sprite *Sprite
}

func NewSpriteNode(texture *ebiten.Image) *SpriteNode {
	sprite := NewSprite(texture)
	return &SpriteNode{sprite: sprite}
}

func (s *SpriteNode) GetSprite() *Sprite {
	return s.sprite
}

func (s *SpriteNode) SetSprite(sprite *Sprite) {
	s.sprite = sprite
}

func (s *SpriteNode) Update() {
	s.updateCurrent()
	for _, child := range s.children {
		child.Update()
	}
}

func (s *SpriteNode) SetPivotToCenter() {
	rect := s.GetSprite().GetTextureRect()
	x, y := rect.GetCenter()
	s.SetPivot(x, y)
}
func (s *SpriteNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM = s.updateGeoM(op.GeoM)
	s.sprite.Draw(target, op)
	for _, child := range s.children {
		child.Draw(target, op)
	}
}
