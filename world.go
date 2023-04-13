package airwar2d

import (
	Assets "github.com/LuigiVanacore/AirWars2D/Assets"
	historia "github.com/LuigiVanacore/AirWars2D/historia_engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	scene  *historia.Scene
	player *Player
	//sprite *historia.SpriteNode
	//desert  *historia.SpriteNode
	//camera  *historia.Camera
	//meteors []*Meteor
}

func NewWorld() *World {
	w := &World{}
	//w.camera = historia.NewCamera(WindowWidth, WindowHeight, WindowWidth/2, WindowHeight/2, 0, 1)
	//w.player = NewPlayer()
	//spriteNode := historia_engine.NewSpriteNode(Assets.ResourceManager().GetTexture(Assets.Ship))
	//spriteNode.SetPivotToCenter()
	entitySprite := historia.NewEntitySprite(historia.NewSprite(Assets.ResourceManager().GetTexture(Assets.Ship)))
	entitySprite.SetPivotToCenter()
	sceneNode := historia.NewSceneNode(entitySprite)
	player := NewPlayer()
	player.SetPosition(150, 150)
	playerNode := historia.NewSceneNode(player)
	playerNode.AttachChild(sceneNode)
	w.scene = historia.NewScene()
	w.scene.AddChild(playerNode, 0)
	//spriteNode := historia.NewSpriteNode(Assets.ResourceManager().GetTexture(Assets.Ship))
	//spriteNode.SetPivotToCenter()
	//spriteNode.SetPosition(200, 200)
	//w.desert = historia.NewSpriteNode(Assets.ResourceManager().GetTexture(Assets.Desert))
	//w.camera.Follow(w.player)
	//w.meteors = append(w.meteors, NewMeteor(big))
	//g.world.AttachChild(g.desert)
	//w.meteors[0].SetPivotToCenter()
	//w.meteors[0].SetPosition(50, 80)
	//c := historia.NewCollisionShape()
	//c.SetDebug(true)
	//w.meteors[0].AttachChild(c)
	//w.AttachChild(w.player)
	//w.AttachChild(spriteNode)
	//w.sprite = spriteNode
	//w.AttachChild(w.meteors[0])
	//w.AttachChild(c)
	return w
}

var rotation int

func (w *World) Update() {
	//rotation += 1
	//w.sprite.SetRotation(rotation)
	//w.correctPosition()
	w.scene.Update()
}

func (w *World) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	w.scene.Draw(target, op)
}

//
//func (w *World) GetPlayer() *Player {
//	return w.player
//}
//
//func (w *World) correctPosition() {
//	playerPosition := w.player.GetTransform().GetPosition()
//	if playerPosition.X < 0 {
//		playerPosition.SetPosition(WindowWidth, WindowHeight-playerPosition.Y)
//	} else if playerPosition.X > WindowWidth {
//		playerPosition.SetPosition(0, WindowHeight-playerPosition.Y)
//	}
//	if playerPosition.Y < 0 {
//		playerPosition.Y = WindowHeight
//	} else if playerPosition.Y > WindowHeight {
//		playerPosition.Y = 0
//	}
//
//	w.player.GetTransform().SetPosition(playerPosition.X, playerPosition.Y)
//}
