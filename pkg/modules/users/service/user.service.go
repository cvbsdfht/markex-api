package service

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/markex-api/pkg/core"
	"github.com/markex-api/pkg/modules"
	userModel "github.com/markex-api/pkg/modules/users/model"
	"github.com/markex-api/pkg/tools/errs"
	"github.com/markex-api/pkg/tools/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Port
type IUserService interface {
	GetUserList() (*[]userModel.User, error)
	GetUserById(id string) (*userModel.User, error)
	Create(user *userModel.User) (*userModel.User, error)
	Update(user *userModel.User) (*userModel.User, error)
	Delete(id string) (*userModel.User, error)
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrNoContent(err)
		}

		return nil, errs.ErrUnexpected(err)
	}

	return users, nil
}

func (s *userService) GetUserById(id string) (*userModel.User, error) {
	Oid := utils.ToObjectID(id)

	user, err := s.repo.UserRepository.GetById(Oid)
	if err != nil {
		s.core.Logger.Error(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrNoContent(err)
		}

		return nil, errs.ErrUnexpected(err)
	}

	return user, nil
}

func (s *userService) Create(user *userModel.User) (*userModel.User, error) {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required),
	)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	now := time.Now()
	request := &userModel.User{
		Id:          primitive.NewObjectID(),
		Email:       user.Email,
		Status:      "registered",
		CreatedDate: now,
		UpdatedDate: now,
	}

	if user.Firstname != "" {
		request.Firstname = user.Firstname
	}
	if user.Lastname != "" {
		request.Lastname = user.Lastname
	}
	if user.Tel != "" {
		request.Tel = user.Tel
	}
	if user.BirthDate != nil {
		request.BirthDate = user.BirthDate
	}

	result, err := s.repo.UserRepository.Upsert(request)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrNotAcceptable(err)
	}

	return result, nil
}

func (s *userService) Update(user *userModel.User) (*userModel.User, error) {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Id, validation.Required),
	)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	now := time.Now()
	request := &userModel.User{
		Id:          user.Id,
		UpdatedDate: now,
	}

	if user.Firstname != "" {
		request.Firstname = user.Firstname
	}
	if user.Lastname != "" {
		request.Lastname = user.Lastname
	}
	if user.Tel != "" {
		request.Tel = user.Tel
	}
	if user.BirthDate != nil {
		request.BirthDate = user.BirthDate
	}

	result, err := s.repo.UserRepository.Upsert(request)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrNotAcceptable(err)
	}

	return result, nil
}

func (s *userService) Delete(id string) (*userModel.User, error) {
	Oid := utils.ToObjectID(id)
	now := time.Now()
	request := &userModel.User{
		Id:          Oid,
		UpdatedDate: now,
		Status:      "closing",
	}

	result, err := s.repo.UserRepository.Upsert(request)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrNotAcceptable(err)
	}

	return result, nil
}
