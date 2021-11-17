package game

type IService interface {
	createGame()
}

type Serivce struct {
}

func (s *Serivce) createGame() *Game {
	return &Game{}
}
