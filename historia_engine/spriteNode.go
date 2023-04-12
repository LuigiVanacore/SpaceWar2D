package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteNode struct {
	*SceneNode
	sprite *Sprite
}

func NewSpriteNode(texture *ebiten.Image) *SpriteNode {
	sprite := NewSprite(texture)
	spriteNode := &SpriteNode{sprite: sprite}
	spriteNode.SceneNode = NewSceneNode(spriteNode)
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
	s.SetPivot(x/2, y/2)
}

func (s *SpriteNode) drawCurrent(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.sprite.Draw(target, op)
}
