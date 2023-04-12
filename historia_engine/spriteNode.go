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
	spriteNode := &SpriteNode{sprite: sprite}
	return spriteNode
}

func (s *SpriteNode) GetSprite() *Sprite {
	return s.sprite
}

func (s *SpriteNode) SetSprite(sprite *Sprite) {
	s.sprite = sprite
}

func (s *SpriteNode) SetPivotToCenter() {
	rect := s.GetSprite().GetTextureRect()
	x, y := rect.GetCenter()
	s.GetTransform().SetPivot(x, y)
}

func (s *SpriteNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	localOp := &ebiten.DrawImageOptions{}
	localOp.GeoM = s.updateGeoM(localOp.GeoM)
	//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
	s.sprite.Draw(target, localOp)
	for _, child := range s.children {
		child.Draw(target, localOp)
	}
}
