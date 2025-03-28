package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	UserRole   Role = "user"
	AdminRole  Role = "admin"
	MasterRole Role = "master"
)

type User struct {
	ID        uint       `json:"user_id" gorm:"primaryKey"`
	Username  string     `json:"user_name" gorm:"size:255;not null"`
	Email     string     `json:"user_email" gorm:"size:255;unique;not null"`
	Password  string     `json:"user_password,omitempty" gorm:"size:255;not null"`
	Role      Role       `json:"user_role" gorm:"type:varchar(10);default:user;not null"`
	CreatedAt time.Time  `json:"user_created_at"`
	UpdatedAt time.Time  `json:"user_updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
	Tickets   []Ticket   `json:"-" gorm:"foreignKey:AuthorID"`
}

// BeforeSave hashs the password before saving
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Only hash if password was changed
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// ComparePassword compares a hashed password with a plain text one
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// ResponseUser is the data structure for user responses to avoid returning sensitive data
type ResponseUser struct {
	Username  string    `json:"user_name"`
	Email     string    `json:"user_email"`
	Role      Role      `json:"user_role"`
	CreatedAt time.Time `json:"user_created_at,omitempty"`
}

// ToResponse converts a User to a ResponseUser
func (u *User) ToResponse(includeCreatedAt bool) ResponseUser {
	response := ResponseUser{
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}

	if includeCreatedAt {
		response.CreatedAt = u.CreatedAt
	}

	return response
}
