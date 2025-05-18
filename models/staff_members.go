package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type StaffMembers struct {
	ID           uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string        `json:"name"`
	EmployeeID   string        `json:"employee_id" gorm:"uniqueIndex;not null"`
	DateOfBirth  time.Time     `json:"date_of_birth"`
	Sex          Sex           `json:"sex" gorm:"type:varchar(20)"`
	UserID       uuid.UUID     `json:"user_id"`
	User         User          `gorm:"foreignKey:UserID"`
	FamilyHeadID *uuid.UUID    `gorm:"family_head"`
	FamilyHead   *StaffMembers `gorm:"foreignKey:FamilyHeadID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *StaffMembers) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
