package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionShape struct {
	SceneNode
	debug bool
}

func NewCollisionShape() *CollisionShape {
	collisionShape := &CollisionShape{}
	return collisionShape
}

func (c *CollisionShape) drawCurrent(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if c.debug {
		//vector.DrawFilledCircle(target, float32(c.Transform.position.X), float32(c.Transform.position.Y), 15, color.White, false)
		//vector.StrokeCircle(target, float32(c.GetTransform().position.X), float32(c.GetTransform().position.Y), 40, 2, color.White, false)

	}
}

func (c *CollisionShape) SetDebug(isDebug bool) {
	c.debug = isDebug
}
