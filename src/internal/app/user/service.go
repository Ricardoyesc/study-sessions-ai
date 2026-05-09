package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"sai-server/internal/domain"
)

type Service struct {
	db        *gorm.DB
	jwtSecret string
}

func NewService(db *gorm.DB, jwtSecret string) *Service {
	return &Service{db: db, jwtSecret: jwtSecret}
}

func (s *Service) Register(ctx context.Context, email, password string) (*domain.User, string, error) {
	existing := domain.User{}
	if err := s.db.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, "", errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := domain.User{
		Email:        email,
		PasswordHash: string(hash),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, "", err
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	user := domain.User{}
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (s *Service) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	user := domain.User{}
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID string, theta, uncertainty *float64, cluster *string) error {
	return s.db.Model(&domain.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"estimated_theta":   theta,
		"theta_uncertainty": uncertainty,
		"cluster":          cluster,
	}).Error
}

func (s *Service) generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
