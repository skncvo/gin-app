package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skncvo/gin-app/config"
	"github.com/skncvo/gin-app/models"
)

// context ? engine
// 호출할 때 매개변수 생략?
// Create User (Post /users)
func CreateUser(c *gin.Context) { //*gin.Context : HTTP 요청 1건에 대한 모든 정보
	var user models.User
	// c.ShouldBindJSON -> 요청 body 파싱
	if err := c.ShouldBindJSON(&user); err != nil {
		// c.JSON -> 응답 보내기
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Read all (Get /users)
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Read one (Get /users/:id)
func GetUserByID(c *gin.Context) {
	var user models.User
	// URL 파라미터 추출
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update (PUT /users/:id)
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.Name = updateData.Name
	user.Email = updateData.Email

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted!"})
}
