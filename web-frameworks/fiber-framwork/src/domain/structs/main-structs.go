package structs

import (
	"fiberframework/src/adapters/validation"
	"fiberframework/src/domain/interfaces"

	"github.com/gofiber/fiber/v2"
)

type ControllerConfig struct {
	App         *fiber.App
	DataService interfaces.IDataService
	Validator   *validation.Validation
}
