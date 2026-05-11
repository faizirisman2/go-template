package handler

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

// NewHealthHandler initializes health handler instance.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Ping godoc
//
// @Summary Health check endpoint
// @Description Check application availability and service status
// @Tags Health
// @Accept json
// @Produce json
//
// @Success 200 {object} map[string]interface{}
//
// @Router /v1/health [get]
func (h *HealthHandler) Ping(c *fiber.Ctx) error {

	result := fiber.Map{
		"message": "pong",
		"status":  "ok",
	}
	return c.JSON(result)
}
