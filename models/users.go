package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserType string

const (
	Admin       UserType = "admin"
	Student     UserType = "student"
	Staff       UserType = "staff"
	StaffFamily UserType = "staff_family"
)

type UserStatus int

const (
	Active    UserStatus = 1
	Inactive  UserStatus = 2
	Suspended UserStatus = 3
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Password  string     `json:"-"` // Omit from JSON responses
	UserType  UserType   `json:"user_type" gorm:"type:varchar(20)"`
	Status    UserStatus `json:"status" gorm:"type:int"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate hook to hash
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
