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
	UpdateChildren()
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
	GetTransform() *Transform
	updateGeoM(geoM ebiten.GeoM) ebiten.GeoM
	// AttachParent GetWorldPosition() math2d.Vector2D
	//GetWorldTransform() ebiten.GeoM
	AttachParent(node Node)
	Delete()
}
