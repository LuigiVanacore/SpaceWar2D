package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Node interface {
	AttachChild(child Node)
	DetachChild(node Node) bool
	// Update GetTransform() ebiten.GeoM
	//SetTransform(transform ebiten.GeoM)
	Update()
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
	GetPosition() (float64, float64)
	SetPosition(x, y float64)
	// AttachParent GetWorldPosition() math2d.Vector2D
	//GetWorldTransform() ebiten.GeoM
	AttachParent(node Node)
	Delete()
}
