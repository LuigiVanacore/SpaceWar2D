package historia_engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type SceneNode struct {
	Transform
	children []Node
	parent   Node
}

func NewSceneNode() *SceneNode {
	return &SceneNode{}
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

func (s *SceneNode) Update() {
	s.updateCurrent()
	s.updateChildren()
}

func (s *SceneNode) updateCurrent() {

}

func (s *SceneNode) updateChildren() {
	for _, child := range s.children {
		child.Update()
	}
}

func (s *SceneNode) updateGeoM(geoM ebiten.GeoM) ebiten.GeoM {
	geoM.Translate(-s.pivot.X, -s.pivot.Y)
	geoM.Rotate(float64(s.rotation%360) * 2 * math.Pi / 360)
	geoM.Translate(s.position.X, s.position.Y)
	return geoM
}
func (s *SceneNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM = s.updateGeoM(op.GeoM)
	for _, child := range s.children {
		child.Draw(target, op)
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
