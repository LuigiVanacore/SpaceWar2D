package historia_engine

import (
	"github.com/LuigiVanacore/AirWars2D/historia_engine/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Node interface {
	AttachChild(child Node)
	DetachChild(node Node) bool
	// Update GetTransform() ebiten.GeoM
	//SetTransform(transform ebiten.GeoM)
	Update()
	UpdateCurrent()
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
	drawCurrent(target *ebiten.Image, op *ebiten.DrawImageOptions)
	GetPosition() math2d.Vector2D
	SetPosition(x, y float64)
	GetTransform() Transform
	// AttachParent GetWorldPosition() math2d.Vector2D
	//GetWorldTransform() ebiten.GeoM
	AttachParent(node Node)
	Delete()
}
