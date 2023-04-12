package airwar2d

import (
	"github.com/LuigiVanacore/AirWars2D/Assets"
	historia "github.com/LuigiVanacore/AirWars2D/historia_engine"
)

const (
	small = iota
	medium
	big
)

type Meteor struct {
	historia.SpriteNode
	_type int
}

func NewMeteor(_type int) *Meteor {
	var spriteNode *historia.SpriteNode
	switch _type {
	case small:
		break
	case medium:
		break
	case big:
		spriteNode = historia.NewSpriteNode(Assets.ResourceManager().GetTexture(Assets.BigMeteor))
		break
	}
	return &Meteor{SpriteNode: *spriteNode}
}
