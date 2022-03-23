package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wryonik/appointment/models"
)

type Reports struct {
	DoctorId       uint
	PatientId      uint
	HospitalId     uint
	ReportFiles    string
	Date           time.Time
	CombinedPdfUrl string
}

type CreateReportInput struct {
	DoctorId    uint      `json:"doctor_id" binding:"required"`
	PatientId   uint      `json:"patient_id" binding:"required"`
	HospitalId  uint      `json:"hospital_id" binding:"required"`
	ReportFiles string    `json:"report_files" binding:"required"`
	Date        time.Time `json:"date_time" binding:"required"`
}

type UpdateReportInput struct {
	DoctorId    uint      `json:"doctor_id" binding:"required"`
	PatientId   uint      `json:"patient_id" binding:"required"`
	HospitalId  uint      `json:"hospital_id" binding:"required"`
	ReportFiles string    `json:"report_files" binding:"required"`
	Date        time.Time `json:"date_time" binding:"required"`
}

// GET /reports
// Find all reports
func FindReports(c *gin.Context) {
	var reports []Reports
	models.DB.Find(&reports)

	c.JSON(http.StatusOK, gin.H{"data": reports})
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

	// Create report
	report := Reports{DoctorId: input.DoctorId, PatientId: input.PatientId, HospitalId: input.HospitalId, Date: input.Date, ReportFiles: input.ReportFiles}
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
	var input UpdateReportInput
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
