package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"zayyid-go/domain/testimoni/feature"
)

type integrationMenuHandler struct {
	feature         *feature.TestimoniFeature
	isRequestLogged bool
}

func NewTestimoniHandler(feature *feature.TestimoniFeature, isRequestLogged bool) UserHandlerInterface {
	return &integrationMenuHandler{
		feature:         feature,
		isRequestLogged: isRequestLogged,
	}
}

func (h integrationMenuHandler) Ping(c *fiber.Ctx) error {
	response := http.StatusText(http.StatusOK)

	h.feature.Ping(c.Context())

	return c.Status(http.StatusOK).JSON(response)
}
