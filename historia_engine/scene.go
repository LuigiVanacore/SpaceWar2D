package historia_engine

import "github.com/hajimehoshi/ebiten/v2"

type Scene struct {
	root   *SceneNode
	layers map[int][]*SceneNode
}

func NewScene() *Scene {
	root := NewSceneNode(nil)
	return &Scene{root: root, layers: make(map[int][]*SceneNode)}
}

func (s *Scene) AddChild(sceneNode *SceneNode, layer int) {
	s.root.AttachChild(sceneNode)
	s.addChildToLayers(sceneNode, layer)
}

func (s *Scene) addChildToLayers(sceneNode *SceneNode, layer int) {
	if _, ok := sceneNode.entity.(Drawable); ok {
		s.layers[layer] = append(s.layers[layer], sceneNode)
	}
	for _, child := range sceneNode.children {
		s.addChildToLayers(child, layer)
	}
}

func (s *Scene) Update() {
	s.root.Update()
	s.root.updateTransform(Transform{})
}

func (s *Scene) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	for _, layer := range s.layers {
		for _, sceneNode := range layer {
			if entity, ok := sceneNode.entity.(Drawable); ok {
				entity.Draw(target, op)
			}
			//vector.DrawFilledCircle(target, float32(s.Transform.pivot.X), float32(s.Transform.pivot.Y), 3, s.color, false)
		}
	}
}
