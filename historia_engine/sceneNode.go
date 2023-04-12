package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

type SceneNode struct {
	Transform
	children []Node
	parent   Node
	color    color.Color
}

func NewSceneNode() *SceneNode {
	sceneNode := &SceneNode{color: colornames.White}
	return sceneNode
}

func (s *SceneNode) SetDebugColor(color color.Color) {
	s.color = color
}

func (s *SceneNode) AttachChild(child Node) {
	child.AttachParent(s)

	s.children = append(s.children, child)
}

func (s *SceneNode) DetachChild(node Node) bool {
	for i, child := range s.children {
		if child == node {
			s.children[i] = s.children[len(s.children)-1]
			s.children = s.children[:len(s.children)-1]
			return true
		}
	}
	return false
}

func (s *SceneNode) AttachParent(node Node) {
	s.parent = node
}

func (s *SceneNode) GetChildren() []Node {
	return s.children
}

func (s *SceneNode) GetTransform() *Transform {
	return &s.Transform
}

func (s *SceneNode) Update() {
	s.UpdateChildren()
}

func (s *SceneNode) UpdateChildren() {
	for _, child := range s.children {
		child.Update()
	}
}

func (s *SceneNode) updateGeoM(geoM ebiten.GeoM) ebiten.GeoM {
	position := s.GetPosition()
	pivot := s.GetPivot()
	rotation := s.GetRotation()
	if s.parent != nil {
		parentPosition := s.parent.GetTransform().GetPosition()
		parentPivot := s.parent.GetTransform().GetPivot()
		parentRotation := s.parent.GetTransform().GetRotation()
		position.Add(parentPosition)
		pivot.Add(parentPivot)
		rotation += parentRotation
	}
	geoM.Translate(-pivot.X, -pivot.Y)
	geoM.Rotate(float64(rotation%360) * 2 * math.Pi / 360)
	geoM.Translate(position.X, position.Y)
	return geoM
}

func (s *SceneNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	localOp := &ebiten.DrawImageOptions{}
	localOp.GeoM = s.updateGeoM(localOp.GeoM)
	//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
	for _, child := range s.children {
		child.Draw(target, localOp)
	}
}

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
