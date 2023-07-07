package services

import (
	"fmt"
	"go-fiber-app/config"
	"go-fiber-app/middleware"
	"go-fiber-app/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(users []models.User) ([]byte, error) {
	config.DB.Raw("SELECT * FROM users").Scan(&users)
	GetUsers, err := middleware.Submit(users, "เรียกดูข้อมูลพนักงาน")
	if err != nil {
		fmt.Println("Error:", err)
	}
	return GetUsers, nil
}

func GetUser(user *models.User, id string) ([]byte, error) {
	config.DB.Raw("SELECT * FROM users WHERE Id = ? ", id).Scan(&user)
	// config.DB.Find(&user , id)
	GetUser, err := middleware.Submit(user, "เรียกดูข้อมูลพนักงาน")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return GetUser, nil
}

func SaveUser(user *models.User) ([]byte, error) {
	config.DB.Save(&user)

	SaveUser, err := middleware.Submit(user.ID, "เพิ่มพนักงานสำเร็จแล้ว")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return SaveUser, nil
}

func UpdateUser(id string, updatedUser *models.User) (*models.User, error) {
	user := new(models.User)
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if user.Email == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "User Not Found")
	}

	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email

	result = config.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func DeleteUser(id string) error {
	user := new(models.User)
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	if user.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "User not available")
	}

	result = config.DB.Unscoped().Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
