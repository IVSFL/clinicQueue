package controllers

import(
	"net/http"
	"clinicQueue/models"
	"clinicQueue/config"

	"github.com/gin-gonic/gin"
)

func CreateAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать администратора"})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

func GetAdmin(c *gin.Context){
	id := c.Param("id")
	var admin models.Admin

	if err := config.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Админ не найден"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func UpdateAdmin(c *gin.Context) {
	id := c.Param("id")
	var admin models.Admin

	if err := config.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Админ не найден"})
		return
	}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&admin)
	c.JSON(http.StatusOK, admin)
}