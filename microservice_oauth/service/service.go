package service

import (
	"github.com/diegoclair/microservice_oauth/domain/contract"
	"github.com/diegoclair/microservice_oauth/service/httprest"
)

// Service holds the domain service repositories
type Service struct {
	db contract.RepoManager
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	AccessTokenService(svc *Service) contract.AccessTokenService
	UserAPIService() contract.UserAPIService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) AccessTokenService(svc *Service) contract.AccessTokenService {
	return newAccessTokenService(svc, s.UserAPIService())
}

func (s *serviceManager) UserAPIService() contract.UserAPIService {
	return httprest.NewUserAPI()
}
