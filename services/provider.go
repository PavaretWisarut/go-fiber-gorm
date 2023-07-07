package services

import (
	"go-fiber-app/config"
	"go-fiber-app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(model *models.User) (uint, error) {
	user := new(models.User)

	// Check if email already exists
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", model.Username).Count(&count)

	if count > 0 {
		return 0, fiber.NewError(fiber.StatusConflict, "username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err.Error())
	}

	user.Username = model.Username
	user.Password = string(hash)
	user.FirstName = model.FirstName
	user.LastName = model.LastName
	user.Email = model.Email

	result := config.DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	userID := user.ID

	return userID, nil
}

func ValidateUser(model *models.Login) (*models.User, error) {
	user := new(models.User)

	rows := config.DB.Raw("SELECT * FROM users WHERE username = ?", model.Username).Scan(&user)
	if rows.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusConflict, "ไม่พบข้อมูล user ในฐานระบบ")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(model.Password))

	if err != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "รหัสผ่านของคุณไม่ถูกต้อง")
	}

	return user, nil
}
