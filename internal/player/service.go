package player

type Service struct {
}

type IService interface {
}

func (s *Service) createPlayer() *Player {
	return &Player{}
}
