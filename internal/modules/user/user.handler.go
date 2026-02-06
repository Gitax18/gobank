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
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "bad request"})
	}

	err = h.s.CreateUser(*user.Name, *user.Number, *user.AccountNumber, *user.Balance); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "error creating user", "err": err.Error()})
	}

	return nil
}

func(h *Handler) PUTUpdateUser(context fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Id should not be empty"})
	}

	type updates struct {
		Name *string `json:"name"`
		Number *int `json:"number"`
	}

	var u updates

	err = context.Bind().JSON(&u); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "improper data", "data": u})
	}
	
	err = h.s.UpdateUser(id, u.Name, u.Number); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "error occured while updating user", "err": err.Error()})
	}	

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user updated successfully",
	})

	return nil
}

func(h *Handler) DELETEUser(context fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Id should not be empty"})
	}
	
	err = h.s.DeleteUser(id); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error deleting user", "err": err.Error()})
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})

	return nil
}

func(h *Handler) GETUser(context fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")); if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Improper ID"})
		return nil
	}

	var user *User

	user, err = h.s.ReadUser(id); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting user",
			"err": err.Error(),
		})
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user found successfully",
		"data": user,
	})

	return nil
}