package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "github.com/rs/zerolog/log"
	// "github.com/davecgh/go-spew/spew"
)


func ProductIdHeader(c *fiber.Ctx) error {
	ProductId := c.Get("X-Product-ID")

	// Jika ProductId tidak ada di header, generate UUID baru
	if ProductId == "" {
		ProductId = generateUUID() // Asumsi: Anda memiliki fungsi generateUUID() yang membuat UUID baru.
	}

	// Set ProductId ke dalam Locals
	c.Locals("product_id", ProductId)
	return c.Next()
}

func generateUUID() string {
	return uuid.New().String() // Menggunakan paket github.com/google/uuid untuk generate UUID
}
