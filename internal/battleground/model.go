package battleground

import (
	player "server/internal/player"
)

type IBattleground interface {
	create_battleground(x float32, y float32)
}

type Battleground struct {
	x_length float32
	y_length float32
	player_a player.Player
	player_b player.Player
	//TODO: bullets 
}

func (s *Battleground) create_battleground(x float32, y float32) {
	s.x_length = x;
	s.y_length = y;
}

// func (s *Battleground) get_



