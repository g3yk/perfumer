package handler

import (
	"log"
	"strconv"
	"time"

	"API/database"
	"API/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserData struct {
	model.BaseUser
	Password *string `json:"-"`
}

type CreateUserResponse struct {
	HttpResponse
	Data CreateUserData `json:"data"`
}

type UpdateUserInput struct {
	Name string `json:"name" example:"John Doe"`
}

type PasswordInput struct {
	Password string `json:"password" example:"user1234"`
}

func hashPassword(password string) (string, error) {
	t1 := time.Now()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	elapsed := time.Since(t1)
	log.Println("Elapsed hashing time:", elapsed)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// GetUser get a user
//
//	@Summary		Get a user
//	@Description	Get a user by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"	default(5)
//	@Success		200	{object}	HttpResponse{data=model.UserResponse}
//	@Failure		404	{object}	HttpResponse{}
//	@Failure		500	{object}	HttpResponse{}
//	@Router			/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// GetDeletedUsers get all deleted users
//
//	@Summary		Get all deleted users
//	@Description	Get all deleted users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	HttpResponse{data=[]model.UserResponse}
//	@Router			/user/deleted [get]
func GetDeletedUsers(c *fiber.Ctx) error {
	log.Println("Get deleted users")
	db := database.DB
	var users []model.User
	// db.Unscoped().Find(&users, "deleted_at IS NOT NULL")
	db.Unscoped().Find(&users, "deleted_at IS NOT NULL")

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted users found",
		"data":    users,
	})
}

// CreateUser new user
//
//	@Summary		Create a new user                                                              j
//	@Description	Create a new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			createUserInput	body		model.BaseUser	true	"User credentials"
//	@Success		200				{object}	CreateUserResponse
//	@Failure		400				{object}	HttpResponse{}
//	@Failure		500				{object}	HttpResponse{}
//	@Router			/user [post]
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request body", "data": err.Error()})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err.Error()})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err.Error()})
	}

	newUser := CreateUserData{
		BaseUser: model.BaseUser{
			Email:    user.Email,
			Username: user.Username,
			Name:     user.Name,
		},
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// UpdateUser update user
//
//	@Summary		Update a user
//	@Description	Update a user by ID
//	@Tags			user
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int				true	"User ID"	default(5)
//	@Param			updateUserInput	body		UpdateUserInput	true	"User credentials"
//	@Success		200				{object}	HttpResponse{data=model.UserResponse}
//	@Failure		401				{object}	HttpResponse{}
//	@Failure		500				{object}	HttpResponse{}
//	@Router			/user/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := database.DB
	var user model.User

	db.First(&user, id)
	user.Name = uui.Name
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

// DeleteUser delete user
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			user
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int				true	"User ID"
//	@Param			passwordInput	body		PasswordInput	true	"User credentials"
//	@Success		200				{object}	HttpResponse{data=nil}
//	@Failure		401				{object}	HttpResponse{}
//	@Failure		500				{object}	HttpResponse{}
//	@Router			/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err.Error(),
		})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	if !validUser(id, pi.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Not valid password",
			"data":    nil,
		})
	}

	db := database.DB
	var user model.User

	db.First(&user, id)

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}

// PermanentDeleteUser delete user
//
//	@Summary		Delete a user
//	@Description	Delete a user by ID
//	@Tags			user
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int				true	"User ID" default(1)
//	@Param			passwordInput	body		PasswordInput	true	"User credentials"
//	@Success		200				{object}	HttpResponse{data=nil}
//	@Failure		400				{object}	HttpResponse{}
//	@Failure		401				{object}	HttpResponse{}
//	@Router			/user/permanent/{id} [delete]
func PermanentDeleteUser(c *fiber.Ctx) error {
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err.Error(),
		})
	}
	id := c.Params("id")

	db := database.DB
	var user model.User

	db.Unscoped().First(&user, id)

	db.Unscoped().Delete(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully deleted",
		"data":    nil,
	})
}
