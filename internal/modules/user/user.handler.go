package user

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	s *Service
}

func(h *Handler) POSTCreateUser(context fiber.Ctx) error {
	var user User
	
	err := context.Bind().JSON(&user); if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "bad request"})
		return err
	}

	err = h.s.CreateUser(*user.Name, *user.Number, *user.AccountNumber, *user.Balance); if err != nil {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "error creating user"})
		return err
	}

	return nil
}

func(h *Handler) PUTUpdateUser(context fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")); if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Id should not be empty"})
		return nil
	}

	type updates struct {
		Name *string `json:"name"`
		Number *int `json:"number"`
	}

	var u updates

	err = context.Bind().JSON(&u); if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "improper data", "data": u})
		return err
	}

	return h.s.UpdateUser(id, u.Name, u.Number)	
}

func(h *Handler) DELETEUser(context fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")); if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Id should not be empty"})
		return nil
	}

	return h.s.DeleteUser(id)
}

func(h *Handler) GETUser(context fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")); if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Improper ID"})
		return nil
	}

	var user *User

	user, err = h.s.ReadUser(id); if err != nil {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting user",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user found successfully",
		"data": user,
	})

	return nil
}