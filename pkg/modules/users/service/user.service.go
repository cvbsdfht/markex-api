package service

import (
	"github.com/markex-api/pkg/core"
	"github.com/markex-api/pkg/core/utils"
	"github.com/markex-api/pkg/modules"
	userModel "github.com/markex-api/pkg/modules/users/model"
)

// Port
type IUserService interface {
	GetUserList() (*[]userModel.User, error)
	GetUserById(id string) (*userModel.User, error)
}

// Adaptor
type userService struct {
	core *core.CoreRegistry
	repo *modules.RepositoryRegistry
}

func NewUserService(c *core.CoreRegistry, r *modules.RepositoryRegistry) IUserService {
	return &userService{core: c, repo: r}
}

func (s *userService) GetUserList() (*[]userModel.User, error) {
	users, err := s.repo.UserRepository.GetAll()
	if err != nil {
		s.core.Logger.Error(err)
		return nil, err
	}

	return users, nil
}

func (s *userService) GetUserById(id string) (*userModel.User, error) {
	Oid := utils.ToObjectID(id)

	user, err := s.repo.UserRepository.GetById(Oid)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, err
	}

	return user, nil
}
