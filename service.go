package main

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

// func (s *Service) FindUser(id uuid.UUID) (*User, error) {
// 	user, err := s.Repository.GetUser(id)

// 	if err != nil {
// 		errMsg := errors.New("user not found")
// 	}
// }
