package user

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	s *Service
}

func(h *Handler) POSTLogoutUser(context fiber.Ctx) error {
		// send the cookie with jwt payload
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Secure = false // because localhost

	context.Cookie(cookie)

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "user logout successfully"})
}

func(h *Handler) POSTLoginUser(context fiber.Ctx) error {

	type UserCreds struct {
		Email 	  string  	`json:"email"`
		Password  string	`json:"password"`
	}
	var uc UserCreds

	err := context.Bind().JSON(&uc); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "bad request"})
	}
	
	user, err := h.s.ReadUserByMail(uc.Email); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "mail not found", "err": err})
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(*user.HashedPassword), []byte(uc.Password)); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "password not found", "err": err})
	}
		
	// sign the JWT with user data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sig": user,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "fail to generate token", "err": err})
	}

	// send the cookie with jwt payload
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = false // because localhost

	context.Cookie(cookie)

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "user logined successfully"})
	
}

func(h *Handler) POSTCreateUser(context fiber.Ctx) error {
	var user User
	
	err := context.Bind().JSON(&user); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "bad request"})
	}

	err = h.s.CreateUser(*user.Email, *user.HashedPassword, *user.Name, *user.Number, *user.AccountNumber, *user.Balance); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "error creating user", "err": err.Error()})
	}
	
	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "user created successfully"})
}

func(h *Handler) PUTUpdateUser(context fiber.Ctx) error {
	// getting id from request local storage (user data was injected through auth middleware)
	t := context.Locals("user").(map[string]any)
	
	id:= int(t["id"].(float64))

	type updates struct {
		Name *string `json:"name"`
		Number *int `json:"number"`
	}

	var u updates

	err := context.Bind().JSON(&u); if err != nil {
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
	// getting id from request local storage (user data was injected through auth middleware)
	t := context.Locals("user").(map[string]any)
	
	id:= int(t["id"].(float64))
	
	err := h.s.DeleteUser(id); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error deleting user", "err": err.Error()})
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})

	return nil
}

func(h *Handler) GETUser(context fiber.Ctx) error {
	// getting id from request local storage (user data was injected through auth middleware)
	t := context.Locals("user").(map[string]any)
	
	id:= int(t["id"].(float64))
	
	var user *User

	user, err := h.s.ReadUser(id); if err != nil {
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