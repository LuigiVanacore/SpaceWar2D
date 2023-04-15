package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

type SceneNode struct {
	entity   Entity
	children []*SceneNode
	parent   *SceneNode
	color    color.Color
}

func NewSceneNode(entity Entity) *SceneNode {
	sceneNode := &SceneNode{entity: entity, color: colornames.White}
	return sceneNode
}

func (s *SceneNode) SetDebugColor(color color.Color) {
	s.color = color
}

func (s *SceneNode) AttachChild(child *SceneNode) {
	child.AttachParent(s)

	s.children = append(s.children, child)
}

func (s *SceneNode) DetachChild(node *SceneNode) bool {
	for i, child := range s.children {
		if child == node {
			s.children[i] = s.children[len(s.children)-1]
			s.children = s.children[:len(s.children)-1]
			return true
		}
	}
	return false
}

func (s *SceneNode) AttachParent(node *SceneNode) {
	s.parent = node
}

func (s *SceneNode) GetChildren() []*SceneNode {
	return s.children
}

func (s *SceneNode) Update() {
	s.UpdateChildren()
	if s.entity != nil {
		s.entity.Update()

	}
}

func (s *SceneNode) UpdateChildren() {
	for _, child := range s.children {
		child.Update()
	}
}

func (s *SceneNode) updateTransform(parent_transform Transform) {
	if entity, ok := s.entity.(Transformable); ok {
		geoM := ebiten.GeoM{}

		transform := entity.GetTransform()
		parentPosition := parent_transform.GetPosition()
		parentRotation := parent_transform.GetRotation()
		position := transform.position.Add(parentPosition)
		pivot := transform.pivot
		rotation := transform.rotation + parentRotation

		geoM.Translate(-pivot.X, -pivot.Y)
		geoM.Rotate(float64(rotation%360) * 2 * math.Pi / 360)
		geoM.Translate(position.X, position.Y)

		transform.SetGeoM(geoM)
		entity.SetTransform(transform)

		parent_transform.SetPosition(position.X, position.Y)
		parent_transform.SetRotation(rotation)
		parent_transform.SetGeoM(geoM)
	}
	for _, child := range s.children {
		child.updateTransform(parent_transform)
	}
}

//func (s *SceneNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
//	localOp := &ebiten.DrawImageOptions{}
//	localOp.GeoM = s.updateGeoM(localOp.GeoM)
//	if Aer, ok := s.entity.(GetSprite); ok {
//		//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
//	for _, child := range s.children {
//		child.Draw(target, localOp)
//	}
//}

//func (s *SceneNode) GetWorldPosition() math2d.Vector2D {
//	transform := s.GetWorldTransform()
//	x, y := transform.Apply(0, 0)
//	return math2d.Vector2D{X: x, Y: y}
//}
//
//func (s *SceneNode) GetWorldTransform() ebiten.GeoM {
//	transform := ebiten.GeoM{}
//
//	for node := Node(s); node != nil; node = s.parent {
//		getTransform := node.GetTransform()
//		transform.Concat(getTransform)
//	}
//
//	return transform
//}
//
//func (s *SceneNode) Translate(x, y float64) {
//	s.transform.Translate(x, y)
//}

func (s *SceneNode) Delete() {
	s.parent.DetachChild(s)
}
