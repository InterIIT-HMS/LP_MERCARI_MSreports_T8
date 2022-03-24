package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wryonik/appointment/models"
)

type Reports struct {
	Doctor         []Doctor
	Patient        Patient
	Hospital       []Hospital
	ReportFiles    string
	Date           time.Time
	CombinedPdfUrl string
}

type Patient struct {
	PatientID uint   `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"not null"`
	NHID      string `gorm:"unique"`
	Gender    string
	Age       uint
}

type CreateReportInput struct {
	DoctorId    uint   `json:"doctor_id" binding:"required"`
	PatientId   uint   `json:"patient_id" binding:"required"`
	HospitalId  uint   `json:"hospital_id" binding:"required"`
	ReportFiles string `json:"report_files" binding:"required"`
	Date        string `json:"date_time" binding:"required"`
}

type Doctor struct {
	DoctorID    uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Degree      string
	Profession  string
	Experience  uint
	PhoneNumber string
	Hospitals   []*Hospital `gorm:"many2many:doctor_hospital;"`
}

type Hospital struct {
	HospitalID  uint `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Address     string
	PhoneNumber string
	Rating      uint
	Doctors     []*Doctor `gorm:"many2many:doctor_hospital;"`
}

// GET /reports
// Find all reports
func FindReports(c *gin.Context) {
	var reports []Reports
	models.DB.Find(&reports)

	c.JSON(http.StatusOK, gin.H{"data": reports})
}

func FindDoctorById(id uint) (Doctor, error) {
	var doctor Doctor
	if err := models.DB.Preload("Hospitals").Where("doctor_id = ?", id).First(&doctor).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}

func FindPatientById(id uint) (Patient, error) {
	var patient Patient
	if err := models.DB.Preload("Hospitals").Where("patient_id = ?", id).First(&patient).Error; err != nil {
		return patient, err
	}

	return patient, nil
}

func FindHospitalById(id uint) (Hospital, error) {
	var hospital Hospital
	if err := models.DB.Preload("Hospitals").Where("hospital_id = ?", id).First(&hospital).Error; err != nil {
		return hospital, err
	}

	return hospital, nil
}

// GET /reports/:id
// Find a report
func FindReport(c *gin.Context) {
	// Get model if exist
	var report Reports
	if err := models.DB.Where("id = ?", c.Param("id")).First(&report).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": report})
}

// POST /reports
// Create new report
func CreateReport(c *gin.Context) {
	// Validate input
	var input CreateReportInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor := make([]Doctor, 1)

	doctor[0], _ = FindDoctorById(input.DoctorId)

	patient, _ := FindPatientById(input.PatientId)

	hospitals := make([]Hospital, 1)
	hospitals[0], _ = FindHospitalById(input.HospitalId)

	// Create report
	report := Reports{Doctor: doctor, Patient: patient, Hospital: hospitals, Date: time.Now(), ReportFiles: input.ReportFiles}
	models.DB.Create(&report)

	c.JSON(http.StatusOK, gin.H{"data": report})
}

// PATCH /reports/:id
// Update a report
func UpdateReport(c *gin.Context) {
	// Get model if exist
	var report Reports
	if err := models.DB.Where("id = ?", c.Param("id")).First(&report).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input CreateReportInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&report).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": report})
}

// DELETE /reports/:id
// Delete a report
func DeleteReport(c *gin.Context) {
	// Get model if exist
	var report Reports
	if err := models.DB.Where("id = ?", c.Param("id")).First(&report).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&report)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
