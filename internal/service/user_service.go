package service

import (
	"errors"

	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/dto"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/entity"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/repository"
	"github.com/ahmadzakyarifin/gin-jwt-auth/utils"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) AuthService {
	return &authService{repo: r}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return errors.New("Email sudah terdaftar")
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPass,
		Role:     "user",
	}

	return s.repo.Create(user)
}

func (s *authService) Login(req *dto.LoginRequest) (string,error){
	user,existing := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return "",errors.New("Email yang anda masukkan salah")
	}
	err := utils.CheckPassword(req.Password,user.Password);
	if err != nil {
		return  "",errors.New("Password yang anda masukkan salah")
	}
	return  utils.GenerateToken(user.ID,user.Role)
}


