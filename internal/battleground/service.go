package battleground

import (
	"server/internal/player"
)

type Service struct {

}

type IService interface {
	create_battleground(x float32, y float32)
}

func (s *Service) create_battleground(
	x float32, 
	y float32, 
	player_a player.Player,
	player_b player.Player) (*Battleground) {
	return &Battleground{x, y, player_a, player_b}
}




