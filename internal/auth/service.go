package auth

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db     *database.DB
	secret []byte
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewService(db *database.DB, secret string) *Service {
	return &Service{
		db:     db,
		secret: []byte(secret),
	}
}

func (s *Service) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (s *Service) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) GenerateToken(user *models.User) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) ValidateToken(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func (s *Service) Login(username, password string) (string, *models.User, error) {

	var user models.User
	err := s.db.QueryRow(
		"SELECT id, username, email, password, role, is_active, last_login, created_at, updated_at FROM users WHERE username = $1",
		username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.LastLogin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	if !user.IsActive {
		return "", nil, errors.New("user account is deactivated")
	}

	if !s.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("invalid username or password")
	}

	_, err = s.db.Exec(
		"UPDATE users SET last_login = $1 WHERE id = $2",
		time.Now(), user.ID)

	if err != nil {

	}

	token, err := s.GenerateToken(&user)
	if err != nil {
		return "", nil, err
	}

	user.Password = ""

	return token, &user, nil
}

func (s *Service) Register(userReg *models.UserRegister) (*models.User, error) {

	hashedPassword, err := s.HashPassword(userReg.Password)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = s.db.QueryRow(
		"INSERT INTO users (username, email, password, role, is_active) VALUES ($1, $2, $3, $4, $5) RETURNING id, username, email, role, is_active, last_login, created_at, updated_at",
		userReg.Username, userReg.Email, hashedPassword, userReg.Role, true).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.IsActive, &user.LastLogin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow(
		"SELECT id, username, email, role, is_active, last_login, created_at, updated_at FROM users WHERE id = $1",
		id).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.IsActive, &user.LastLogin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
