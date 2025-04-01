package repository

import (
	"errors"

	"hcall/api/database"
	"hcall/api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.DB,
	}
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// FindByEmail finds a user by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// FindMaster finds a master user
func (r *UserRepository) FindMaster() (*models.User, error) {
	var user models.User
	result := r.DB.Where("role = ?", models.MasterRole).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("master user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// GetUsers gets all users
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByRole gets users by role
func (r *UserRepository) GetUsersByRole(role models.Role) ([]models.User, error) {
	var users []models.User
	if err := r.DB.Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found with specified role")
	}

	return users, nil
}

// DeleteUser deletes a user by email
func (r *UserRepository) DeleteUser(email string) error {
	return database.ExecuteInTransaction(r.DB, func(tx *gorm.DB) error {
		// Verifique se existe algum ticket associado a este usuário
		var count int64
		if err := tx.Model(&models.Ticket{}).Where("author_email = ?", email).Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return errors.New("cannot delete user with existing tickets")
		}

		result := tx.Where("email = ?", email).Delete(&models.User{})

		if result.RowsAffected == 0 {
			return errors.New("user not found")
		}

		return result.Error
	})
}

// DeleteMaster deletes the master user
func (r *UserRepository) DeleteMaster() error {
	return database.ExecuteInTransaction(r.DB, func(tx *gorm.DB) error {
		// Verifique se existe algum ticket associado ao master
		var master models.User
		if err := tx.Where("role = ?", models.MasterRole).First(&master).Error; err != nil {
			return errors.New("master user not found")
		}

		var count int64
		if err := tx.Model(&models.Ticket{}).Where("author_email = ?", master.Email).Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return errors.New("cannot delete master with existing tickets")
		}

		result := tx.Where("role = ?", models.MasterRole).Delete(&models.User{})

		if result.RowsAffected == 0 {
			return errors.New("master user not found")
		}

		return result.Error
	})
}
