package handler

import (
	"API/database"
	"API/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CreatePerfumeNoteInput struct {
	Note string `json:"note" validate:"required" example:"Amber"`
}

type CreatePerfumeNoteResponse struct {
	PerfumeNoteID uint   `json:"perfume_note_id" validate:"required" example:"1"`
	Note          string `json:"note" validate:"required" example:"Amber"`
}

type CreatePerfumeInput struct {
	Name   string `json:"name"`
	Amount *int   `json:"amount" example:"1"`

	BaseNotes  []uint `json:"baseNotes" example:"1"`
	HeartNotes []uint `json:"heartNotes" example:"1"`
	TopNotes   []uint `json:"topNotes" example:"1"`
}

type CreatePerfumeResponse struct {
	PerfumeID uint      `json:"perfume_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2025-04-23T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-04-23T00:00:00Z"`

	Name   string `json:"name"`
	Amount *int   `json:"amount" example:"1"`

	BaseNotes  []CreatePerfumeNoteInput `json:"baseNotes"`
	HeartNotes []CreatePerfumeNoteInput `json:"heartNotes"`
	TopNotes   []CreatePerfumeNoteInput `json:"topNotes"`
}

type AddPerfumetoOrderInput struct {
	OrderID    uint   `json:"order_id" example:"1"`
	PerfumeIDs []uint `json:"perfume_ids" example:"1"`
}

type GetOrderResponse struct {
	OrderID  uint                    `json:"order_id" example:"1"`
	Perfumes []CreatePerfumeResponse `json:"perfumes"`
}

// CreatePerfurmeNote new PerfumeNote
//
//	@Summary		Create a new PerfumeNote
//	@Description	Create a new PerfumeNote
//	@Tags			perfume
//	@Accept			json
//	@Produce		json
//	@Param			createPerfumeNoteInput	body		CreatePerfumeNoteInput	true	"Perfume Note"
//	@Success		200						{object}	HttpResponse{data=CreatePerfumeNoteResponse}
//	@Failure		500						{object}	HttpResponse{}
//	@Router			/perfume/note [post]
func CreatePerfurmeNote(c *fiber.Ctx) error {
	db := database.DB

	perfumeNote := new(model.PerfumeNote)

	if err := c.BodyParser(perfumeNote); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create Perfume Note",
			"data":    err,
		})
	}

	if err := db.Create(&perfumeNote).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create perfume note", "data": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Perfume Note",
		"data":    perfumeNote,
	})
}

// CreatePerfume new Perfume
//
//	@Summary		Create a new Perfume
//	@Description	Create a new Perfume
//	@Tags			perfume
//	@Accept			json
//	@Produce		json
//	@Param			createPerfumeInput	body		CreatePerfumeInput	true	"Perfume"
//	@Success		200					{object}	HttpResponse{data=CreatePerfumeResponse}
//	@Failure		400					{object}	HttpResponse{}
//	@Failure		500					{object}	HttpResponse{}
//	@Router			/perfume [post]
func CreatePerfume(c *fiber.Ctx) error {
	db := database.DB

	// perfume := new(model.Perfume)
	perfumeInput := new(CreatePerfumeInput)

	if err := c.BodyParser(perfumeInput); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"data":    err,
		})
	}

	perfume := new(model.Perfume)

	perfume.Name = perfumeInput.Name
	perfume.Amount = perfumeInput.Amount

	for _, noteID := range perfumeInput.BaseNotes {
		note := new(model.PerfumeNote)
		res := db.First(&note, noteID)
		if err := res.Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("The baseNote id %d does not exist", noteID), "data": err.Error()})
		}
		perfume.BaseNotes = append(perfume.BaseNotes, *note)
	}

	for _, noteID := range perfumeInput.HeartNotes {
		note := new(model.PerfumeNote)
		res := db.First(&note, noteID)
		if err := res.Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("The heartNote id %d does not exist", noteID), "data": err.Error()})
		}
		perfume.HeartNotes = append(perfume.HeartNotes, *note)
	}

	for _, noteID := range perfumeInput.TopNotes {
		note := new(model.PerfumeNote)
		res := db.First(&note, noteID)
		if err := res.Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("The topNote id %d does not exist", noteID), "data": err.Error()})
		}
		perfume.TopNotes = append(perfume.TopNotes, *note)
	}

	if err := db.Create(&perfume).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create perfume", "data": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Perfume",
		"data":    perfume,
	})
}

// CreateOrder new Order
//
//	@Summary		Create a new Order
//	@Description	Create a new Order
//	@Tags			order
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	HttpResponse{data=object{id=int}}
//	@Failure		500	{object}	HttpResponse{}
//	@Router			/perfume/order [post]
func CreateOrder(c *fiber.Ctx) error {
	db := database.DB

	order := new(model.Order)
	db.Create(&order)

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	user := new(model.User)

	if err := db.First(&user, uid).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User does not exist", "data": err.Error()})
	}

	db.Model(&user).Association("Orders").Append(order)

	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create order", "data": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Order",
		"data":    fiber.Map{"id": order.ID},
	})
}

// AddPerfumetoOrder add perfume to order
//
//	@Summary		Add perfume to order
//	@Description	Add perfume to order
//	@Tags			order
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			addPerfumetoOrderInput	body		AddPerfumetoOrderInput	true	"Perfume"
//	@Success		200						{object}	HttpResponse{data=object{id=int}}
//	@Failure		500						{object}	HttpResponse{}
//	@Router			/perfume/add [post]
func AddPerfumetoOrder(c *fiber.Ctx) error {
	db := database.DB

	orderInput := new(AddPerfumetoOrderInput)
	token := c.Locals("user").(*jwt.Token)

	if err := c.BodyParser(&orderInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid request body", "data": err.Error()})
	}

	order := new(model.Order)
	res := db.First(&order, orderInput.OrderID)
	if err := res.Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("The order id %d does not exist", orderInput.OrderID), "data": err.Error()})
	}

	if !validToken(token, strconv.Itoa(order.UserID)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, perfumeID := range orderInput.PerfumeIDs {
			perfume := new(model.Perfume)
			res := db.First(&perfume, perfumeID)
			if err := res.Error; err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": fmt.Sprintf("The perfume id %d does not exist", perfumeID), "data": err.Error()})
			}
			tx.Model(&order).Association("Perfumes").Append(perfume)
		}

		return nil
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't add perfume to order", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Added perfume to order", "data": order})
}

// GetOrder query order
//
//	@Summary		Get an order
//	@Description	Get an order
//	@Tags			order
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"	default(1)
//	@Success		200	{object}	HttpResponse{data=GetOrderResponse}
//	@Failure		404	{object}	HttpResponse{}
//	@Failure		500	{object}	HttpResponse{}
//	@Router			/perfume/order/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	db := database.DB

	token := c.Locals("user").(*jwt.Token)

	order := new(model.Order)
	res := db.
		Preload("Perfumes.BaseNotes").
		Preload("Perfumes.HeartNotes").
		Preload("Perfumes.TopNotes").
		Find(&order, c.Params("id"))

	if err := res.Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No order found with ID", "data": res})
	}

	if !validToken(token, strconv.Itoa(order.UserID)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Order found", "data": order})
}
