package airwar2d

import (
	historia "github.com/LuigiVanacore/AirWars2D/historia_engine"
)

type World struct {
	*historia.SceneNode
	player *Player
	//desert  *historia.SpriteNode
	//camera  *historia.Camera
	//meteors []*Meteor
}

func NewWorld() *World {
	w := &World{}
	w.SceneNode = historia.NewSceneNode(w)
	//w.camera = historia.NewCamera(WindowWidth, WindowHeight, WindowWidth/2, WindowHeight/2, 0, 1)
	w.player = NewPlayer()
	//w.desert = historia.NewSpriteNode(Assets.ResourceManager().GetTexture(Assets.Desert))
	//w.camera.Follow(w.player)
	//w.meteors = append(w.meteors, NewMeteor(big))
	//g.world.AttachChild(g.desert)
	//w.meteors[0].SetPivotToCenter()
	//w.meteors[0].SetPosition(50, 80)
	c := historia.NewCollisionShape()
	c.SetDebug(true)
	//w.meteors[0].AttachChild(c)
	w.AttachChild(w.player)
	//w.AttachChild(w.meteors[0])
	w.AttachChild(c)
	return w
}

func (w *World) UpdateCurrent() {
	w.correctPosition()
}

func (w *World) GetPlayer() *Player {
	return w.player
}

func (w *World) correctPosition() {
	playerPosition := w.player.GetPosition()
	if playerPosition.X < 0 {
		playerPosition.SetPosition(WindowWidth, WindowHeight-playerPosition.Y)
	} else if playerPosition.X > WindowWidth {
		playerPosition.SetPosition(0, WindowHeight-playerPosition.Y)
	}
	if playerPosition.Y < 0 {
		playerPosition.Y = WindowHeight
	} else if playerPosition.Y > WindowHeight {
		playerPosition.Y = 0
	}

	w.player.SetPosition(playerPosition.X, playerPosition.Y)
}
