package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FactorUpdateRequest struct {
	Publisher string  `json:"publisher" validate:"required"`
	Device    string  `json:"device" validate:"required,device"`
	Country   string  `json:"country" validate:"required,country"`
	Factor    float64 `json:"factor" validate:"required,factor"`
	Domain    string  `json:"domain"`
}

func ValidateBulkFactors(c *fiber.Ctx) error {
	var requests []FactorUpdateRequest
	err := c.BodyParser(&requests)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body. Please ensure it's a valid JSON.",
		})
	}

	const minFactor string = "0.01"
	const maxFactor string = "10.0"

	var errorMessages = map[string]string{
		"country": "Country code must be 2 characters long and should be in the allowed list",
		"device":  "Device should be in the allowed list",
		"factor":  fmt.Sprintf("Factor value not allowed, it should be >= %s and <= %s", minFactor, maxFactor),
	}

	for _, request := range requests {
		err = Validator.Struct(request)
		if err != nil {
			errorResponse := map[string]string{
				"status": "error",
			}
			for _, err := range err.(validator.ValidationErrors) {
				if msg, ok := errorMessages[err.Tag()]; ok {
					errorResponse["message"] = msg
				} else {
					errorResponse["message"] = fmt.Sprintf("%s is mandatory, validation failed", err.Field())
				}
				break
			}
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
		}
	}

	return c.Next()
}
