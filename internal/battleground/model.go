package battleground

import (
	player "server/internal/player"
)

type IBattleground interface {
	createBattleground(x float32, y float32)
}

type Battleground struct {
	Height  float32
	Width   float32
	PlayerA player.Player
	PlayerB player.Player
	//TODO: bullets
}

func (s *Battleground) createBattleground(x float32, y float32) *Battleground {
	return &Battleground{
		Height: x,
		Width:  y,
	}
}

// func (s *Battleground) get_
