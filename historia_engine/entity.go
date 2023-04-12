package historia_engine

//type Entity struct {
//	Transform
//	parent     *Entity
//	components []*Entity
//}
//
//func (e *Entity) AttachComponent(component *Entity) {
//	child.AttachParent(e)
//
//	e.components = append(s.children, child)
//}
//
//func (e *Entity) DetachChild(child Entity) bool {
//	for i, child := range e.children {
//		if child == node {
//			s.children[i] = s.children[len(s.children)-1]
//			s.children = s.children[:len(s.children)-1]
//			return true
//		}
//	}
//	return false
//}
//
//func (s *SceneNode) AttachParent(node Node) {
//	s.parent = node
//}
//
//func (s *SceneNode) GetChildren() []Node {
//	return s.children
//}
//
//func (s *SceneNode) GetTransform() *Transform {
//	return &s.Transform
//}
//
//func (s *SceneNode) Update() {
//	s.node.UpdateCurrent()
//	s.updateChildren()
//}
//
//func (s *SceneNode) UpdateCurrent() {
//
//}
//
//func (s *SceneNode) updateChildren() {
//	for _, child := range s.children {
//		child.Update()
//	}
//}
//
//func (s *SceneNode) updateGeoM(geoM ebiten.GeoM) ebiten.GeoM {
//	transform := s.node.GetTransform()
//
//	geoM.Translate(-transform.pivot.X, -transform.pivot.Y)
//	geoM.Rotate(float64(transform.rotation%360) * 2 * math.Pi / 360)
//	geoM.Translate(transform.position.X, transform.position.Y)
//	return geoM
//}
//func (s *SceneNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
//	localOp := &ebiten.DrawImageOptions{}
//	localOp.GeoM = s.node.updateGeoM(op.GeoM)
//	vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
//	s.node.drawCurrent(target, localOp)
//	for _, child := range s.children {
//		child.Draw(target, localOp)
//	}
//}
