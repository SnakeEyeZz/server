package game

import (
	"server/internal/battleground"
	"server/internal/bullet"
	"server/internal/player"
)

type GameSnapshot struct {
	playerA      player.Player
	playerB      player.Player
	bullet       bullet.Bullet
	battleground battleground.Battleground
}
