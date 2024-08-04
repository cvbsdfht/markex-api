package service

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
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
	GetUserList() (*userModel.UserListResponse, error)
	GetUserById(id string) (*userModel.UserResponse, error)
	Create(user *userModel.UserRequest) (*userModel.UserResponse, error)
	Update(user *userModel.UserRequest) (*userModel.UserResponse, error)
	Closing(id string) (*userModel.UserResponse, error)
	Login(request *userModel.UserLoginRequest) (*userModel.UserLoginResponse, error)
}

// Adaptor
type userService struct {
	core     *core.CoreRegistry
	repo     *modules.RepositoryRegistry
	validate *validator.Validate
}

func NewUserService(c *core.CoreRegistry, r *modules.RepositoryRegistry, v *validator.Validate) IUserService {
	return &userService{
		core:     c,
		repo:     r,
		validate: v,
	}
}

func (s *userService) GetUserList() (*userModel.UserListResponse, error) {
	// get all users
	users, err := s.repo.UserRepository.GetAll()
	if err != nil {
		s.core.Logger.Error(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrNotFound(err)
		}

		return nil, errs.ErrBadRequest(err)
	}

	return &userModel.UserListResponse{
		Status: true,
		Data:   users,
	}, nil
}

func (s *userService) GetUserById(id string) (*userModel.UserResponse, error) {
	// validate
	if err := s.validate.Var(id, "mongodb"); err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	// get user by id
	oid, _ := primitive.ObjectIDFromHex(id)
	user, err := s.repo.UserRepository.GetById(oid)
	if err != nil {
		s.core.Logger.Error(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrNotFound(err)
		}

		return nil, errs.ErrBadRequest(err)
	}

	return &userModel.UserResponse{
		Status: true,
		Data:   user,
	}, nil
}

func (s *userService) Create(req *userModel.UserRequest) (*userModel.UserResponse, error) {
	// vaiidate
	if err := s.validate.Struct(req); err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	// build a user
	now := time.Now()
	user := &userModel.User{
		Id:          primitive.NewObjectID(),
		Email:       req.Email,
		Status:      userModel.USER_STATUS_REGISTERED,
		CreatedDate: now,
		UpdatedDate: now,
	}

	if req.Firstname != "" {
		user.Firstname = req.Firstname
	}
	if req.Lastname != "" {
		user.Lastname = req.Lastname
	}
	if req.Tel != "" {
		user.Tel = req.Tel
	}
	if req.BirthDate != "" {
		birthDate := utils.StringToDateFormat(req.BirthDate, utils.FORMAT_DATETIME)
		user.BirthDate = &birthDate
	}

	// upsert user
	result, err := s.repo.UserRepository.Upsert(user)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrNotAcceptable(err)
	}

	return &userModel.UserResponse{
		Status: true,
		Data:   result,
	}, nil
}

func (s *userService) Update(req *userModel.UserRequest) (*userModel.UserResponse, error) {
	// validate
	if err := s.validate.StructPartial(req, "Id"); err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	// build update user
	oid, _ := primitive.ObjectIDFromHex(req.Id)
	user := &userModel.User{
		Id:          oid,
		UpdatedDate: time.Now(),
	}

	if req.Firstname != "" {
		user.Firstname = req.Firstname
	}
	if req.Lastname != "" {
		user.Lastname = req.Lastname
	}
	if req.Tel != "" {
		user.Tel = req.Tel
	}
	if req.BirthDate != "" {
		birthDate := utils.StringToDateFormat(req.BirthDate, utils.FORMAT_DATE)
		user.BirthDate = &birthDate
	}

	// upsert user
	result, err := s.repo.UserRepository.Upsert(user)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrNotAcceptable(err)
	}

	return &userModel.UserResponse{
		Status: true,
		Data:   result,
	}, nil
}

func (s *userService) Closing(id string) (*userModel.UserResponse, error) {
	// validate
	if err := s.validate.Var(id, "mongodb"); err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	// build closing user
	oid, _ := primitive.ObjectIDFromHex(id)
	user := &userModel.User{
		Id:          oid,
		UpdatedDate: time.Now(),
		Status:      userModel.USER_STATUS_CLOSING,
	}

	// upsert user
	result, err := s.repo.UserRepository.Upsert(user)
	if err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrNotAcceptable(err)
	}

	return &userModel.UserResponse{
		Status: true,
		Data:   result,
	}, nil
}

func (s *userService) Login(request *userModel.UserLoginRequest) (*userModel.UserLoginResponse, error) {
	// validate
	if err := s.validate.Struct(request); err != nil {
		s.core.Logger.Error(err)
		return nil, errs.ErrValidationFailed(err)
	}

	// get user by email
	user, err := s.repo.UserRepository.GetByEmail(request.Email)
	if err != nil {
		s.core.Logger.Error(err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrNotFound(err)
		}

		return nil, errs.ErrBadRequest(err)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"userId": user.Id.Hex(),
		"email":  user.Email,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, errs.ErrBadRequest(err)
	}

	return &userModel.UserLoginResponse{
		Email: user.Email,
		Token: t,
	}, nil
}
