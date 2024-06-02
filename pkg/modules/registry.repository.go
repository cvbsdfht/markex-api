package modules

import userRepo "github.com/markex-api/pkg/modules/users/repository"

type RepositoryRegistry struct {
	UserRepository userRepo.IUserRepository
}
