package controllers

import (
	"go-fiber-app/middleware"
	"go-fiber-app/models"
	"go-fiber-app/services"
	"go-fiber-app/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	models := new(models.User)
	if err := c.BodyParser(models); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	DataId, err := services.Register(models)

	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Username นี้มีในระบบแล้ว กรุณาลองใหม่อีกครั้ง"})
	}

	Register, err := middleware.Registers(int(DataId), "สมัครสมาชิกสำเร็จ")

	if err != nil {
		panic(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send(Register)
}

func Login(c *fiber.Ctx) error {
	models := new(models.Login)
	
	if err := c.BodyParser(models); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	UsernameValidate, err := services.ValidateUser(models)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	
	token, err := utils.GenerateToken(UsernameValidate)

	if err != nil {
		panic(err.Error())
	}

	LoginToken, err := middleware.GenerateToken(token, "ล็อคอินสำเร็จ")

	if err != nil {
		panic(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send(LoginToken)
}
