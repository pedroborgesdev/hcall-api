package services

import (
	"errors"

	"hcall/api/models"
	"hcall/api/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(username, email, password string, role models.Role) error {
	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	// Create the new user
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	// Save the user to the database
	return s.userRepo.CreateUser(user)
}

// GetUsers gets all users
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepo.GetUsers()
}

// GetUsersByRole gets users by role
func (s *UserService) GetUsersByRole(role models.Role) ([]models.User, error) {
	return s.userRepo.GetUsersByRole(role)
}

// GetUserByEmail gets a user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}

// GetUserByEmailAndRole gets a user by email and role
func (s *UserService) GetUserByEmailAndRole(email string, role models.Role) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Role != role {
		return nil, errors.New("user with specified role not found")
	}

	return user, nil
}

// DeleteUser deletes a user by email
func (s *UserService) DeleteUser(email string) error {
	return s.userRepo.DeleteUser(email)
}
