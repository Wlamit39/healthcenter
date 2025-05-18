package controllers

import (
	"crud-go/models"
	"github.com/google/uuid"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStudentByID(c *gin.Context, db *gorm.DB) {
	studentIDStr := c.Param("id")

	// Parse UUID
	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID format"})
		return
	}

	var student models.Students
	if err := db.Preload("User").First(&student, "id = ?", studentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudents(c *gin.Context, db *gorm.DB) {
	var students []models.Students
	query := db.Preload("User")

	if studentIDStr := c.Query("id"); studentIDStr != "" {
		if studentID, err := uuid.Parse(studentIDStr); err == nil {
			query = query.Where("students.id = ?", studentID)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID format"})
			return
		}
	}

	if rollNumber := c.Query("roll_number"); rollNumber != "" {
		query = query.Where("students.roll_number = ?", rollNumber)
	}

	if sex := c.Query("sex"); sex != "" {
		query = query.Where("students.sex = ?", sex)
	}

	if userType := c.Query("user_type"); userType != "" {
		// Join with users table to filter on user_type
		query = query.Joins("JOIN users ON users.id = students.user_id").
			Where("users.user_type = ?", userType)
	}

	if err := query.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching students"})
		return
	}

	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context, db *gorm.DB) {
	var student struct {
		Name        string     `json:"name"`
		DateOfBirth time.Time  `json:"date_of_birth"`
		RollNumber  string     `json:"roll_number"`
		Sex         models.Sex `json:"sex"`

		Email    string            `json:"email"`
		Phone    string            `json:"phone"`
		UserType models.UserType   `json:"user_type"`
		Status   models.UserStatus `json:"status"`
		Password string            `json:"password"`
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       uuid.New(),
		Name:     student.Name,
		Email:    student.Email,
		Phone:    student.Phone,
		Password: student.Password,
		UserType: student.UserType,
		Status:   student.Status,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	studentPayload := models.Students{
		ID:          uuid.New(),
		Name:        student.Name,
		DateOfBirth: student.DateOfBirth,
		RollNumber:  student.RollNumber,
		Sex:         student.Sex,
		UserID:      user.ID,
	}

	if err := db.Create(&studentPayload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"student": student,
		"user":    user,
	})
}
