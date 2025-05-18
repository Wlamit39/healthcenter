package controllers

import (
	"crud-go/models"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStaffMemberByID(c *gin.Context, db *gorm.DB) {
	staffMemberIDStr := c.Param("id")

	// Parse UUID
	staffMemberID, err := uuid.Parse(staffMemberIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID format"})
		return
	}

	var staffMember models.StaffMembers
	if err := db.Preload("User").First(&staffMember, "id = ?", staffMemberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, staffMember)
}

func GetStaffMembers(c *gin.Context, db *gorm.DB) {
	var staffMembers []models.StaffMembers
	query := db.Preload("User")

	// Filter by ID
	if staffMemberIDStr := c.Query("id"); staffMemberIDStr != "" {
		if staffMemberID, err := uuid.Parse(staffMemberIDStr); err == nil {
			query = query.Where("id = ?", staffMemberID) // no alias needed
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff member ID format"})
			return
		}
	}

	// Filter by Employee ID
	if employeeID := c.Query("employee_id"); employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}

	// Filter by Sex
	if sex := c.Query("sex"); sex != "" {
		query = query.Where("sex = ?", sex)
	}

	// Execute the query
	if err := query.Find(&staffMembers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching staff members"})
		return
	}

	c.JSON(http.StatusOK, staffMembers)
}

func CreateStaffMember(c *gin.Context, db *gorm.DB) {
	var staffMember struct {
		Name        string     `json:"name"`
		DateOfBirth time.Time  `json:"date_of_birth"`
		EmployeeID  string     `json:"employee_id"`
		Sex         models.Sex `json:"sex"`

		Email    string            `json:"email"`
		Phone    string            `json:"phone"`
		UserType models.UserType   `json:"user_type"`
		Status   models.UserStatus `json:"status"`
		Password string            `json:"password"`
	}

	if err := c.ShouldBindJSON(&staffMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       uuid.New(),
		Name:     staffMember.Name,
		Email:    staffMember.Email,
		Phone:    staffMember.Phone,
		Password: staffMember.Password,
		UserType: staffMember.UserType,
		Status:   staffMember.Status,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	staffMemberPayload := models.StaffMembers{
		ID:          uuid.New(),
		Name:        staffMember.Name,
		DateOfBirth: staffMember.DateOfBirth,
		EmployeeID:  staffMember.EmployeeID,
		Sex:         staffMember.Sex,
		UserID:      user.ID,
	}

	if err := db.Create(&staffMemberPayload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"staff_member": staffMember,
		"user":         user,
	})
}
