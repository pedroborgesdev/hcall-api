package services

import (
	"errors"
	"time"

	"hcall/api/models"
	"hcall/api/repository"
	"hcall/api/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}

// Register registers a new user
func (s *AuthService) Register(username, email, password string) (*models.User, string, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return nil, "", errors.New("email already exists")
	}

	err = utils.ValidateCredentials(email, password, username)
	if err != nil {
		return nil, "", err
	}

	// Create the new user
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     models.UserRole,
	}

	// Save the user to the database
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// Login logs in a user and returns a JWT token
func (s *AuthService) Login(email, password string) (*models.User, string, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("email aren't registered")
	}

	// Compare passwords
	if err := user.ComparePassword(password); err != nil {
		return nil, "", errors.New("password is incorrect")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// CreateMaster creates a master user if one doesn't exist
func (s *AuthService) CreateMaster(email, password string) (*models.User, string, error) {
	// Check if a master user already exists
	_, err := s.userRepo.FindMaster()
	if err == nil {
		return nil, "", errors.New("master user already exists")
	}

	// Create the master user
	master := &models.User{
		Username:  "Master",
		Email:     email,
		Password:  password,
		Role:      models.MasterRole,
		CreatedAt: time.Now(),
	}

	// Save the master user to the database
	if err := s.userRepo.CreateUser(master); err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(master)
	if err != nil {
		return nil, "", err
	}

	return master, token, nil
}

// DeleteMaster deletes the master user
func (s *AuthService) DeleteMaster(email, password string) error {
	// Find the master user
	master, err := s.userRepo.FindMaster()
	if err != nil {
		return errors.New("master user not found")
	}

	// Verify the email matches
	if master.Email != email {
		return errors.New("invalid master credentials")
	}

	// Compare passwords
	if err := master.ComparePassword(password); err != nil {
		return errors.New("invalid master credentials")
	}

	// Delete the master user
	if err := s.userRepo.DeleteMaster(); err != nil {
		return err
	}

	return nil
}
