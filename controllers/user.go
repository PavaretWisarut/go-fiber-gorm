package controllers

import (
	"fmt"
	"go-fiber-app/models"
	"go-fiber-app/middleware"
	"go-fiber-app/services"
	"github.com/gofiber/fiber/v2"

)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	GetUsers , err :=  services.GetUsers(users)
	if err != nil {
		panic(err.Error())
	}

	return c.Status(fiber.StatusOK).Send(GetUsers)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	GetUsers , err :=  services.GetUser(&user , id)
	if err != nil {
		panic(err.Error())
	}

	return c.Status(fiber.StatusOK).Send(GetUsers)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	SaveUser , err:= services.SaveUser(user)
	if err != nil {
		panic(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send(SaveUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedUser := new(models.User)

	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	user, err := services.UpdateUser(id, updatedUser)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).SendString(fiberErr.Message)
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	UpdateUser, err := middleware.Submit(user.ID, "อัปเดทข้อมูลพนักงานสำเร็จแล้ว")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return c.Status(fiber.StatusCreated).Send(UpdateUser)
}


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.DeleteUser(id)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).SendString(fiberErr.Message)
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	DeleteUser, err := middleware.Submit(id, "ลบข้อมูลพนักงานสำเร็จแล้ว")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return c.Status(fiber.StatusCreated).Send(DeleteUser)
}