package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Floor struct {
	Publisher string  `json:"publisher" validate:"required"`
	Device    string  `json:"device" validate:"required,device"`
	Country   string  `json:"country" validate:"required,country"`
	Domain    string  `json:"domain"`
	Floor     float64 `json:"floor"`
}

func ValidateFloors(c *fiber.Ctx) error {
	body := new(Floor)
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	var errorMessages = map[string]string{
		"country": "Country code must be 2 characters long and should be in the allowed list",
		"device":  "Device should be in the allowed list",
	}

	err = Validator.Struct(body)
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
	return c.Next()
}
