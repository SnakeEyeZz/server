package game

type IGame interface {
	getGameSnapshot()
}

type Game struct {
}

func (s *Game) getGameSnapshot() GameSnapshot {
	var g = GameSnapshot{}
	return g
}

func (s *Game) updateGame(g *GameSnapshot) {

}
