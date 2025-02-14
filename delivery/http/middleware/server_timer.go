package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		duration := stop.Sub(start)
		seconds := duration.Seconds()
		// Do something with response
		c.Append("X-Response-Time", fmt.Sprintf("%v seconds", seconds))
		// return stack error if exist
		return err
	}
}
