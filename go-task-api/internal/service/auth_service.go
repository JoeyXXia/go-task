package service

import (
	"errors"
	"github.com/joey/go-task/internal/model"
	"github.com/joey/go-task/internal/repository"
	"github.com/joey/go-task/pkg/jwt"
)

// AuthService 处理认证和授权逻辑
type AuthService struct {
	userRepo *repository.UserRepository
}

// NewAuthService 创建一个新的认证服务
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

// Register 注册新用户
func (s *AuthService) Register(username, email, password string) (*model.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// 检查邮箱是否已存在
	existingEmail, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingEmail != nil {
		return nil, errors.New("email already exists")
	}

	// 创建新用户
	user := &model.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *AuthService) Login(username, password string) (string, error) {
	// 查找用户
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid username or password")
	}

	// 验证密码
	if err := user.ComparePassword(password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// 生成JWT令牌
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken 验证JWT令牌
func (s *AuthService) ValidateToken(tokenString string) (uint, error) {
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	// 检查用户是否存在
	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("user not found")
	}

	return claims.UserID, nil
}
