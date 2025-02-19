package context

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
	"zayyid-go/domain/shared/helper/constant"
	"zayyid-go/domain/shared/model"
)

func CreateContext() context.Context {
	return context.Background()
}

func SetValueToContext(ctx context.Context, c *fiber.Ctx) context.Context {
	Token := c.Get("Authorization")
	UserId, ok := c.Locals("x-user-id").(string)
	if !ok {
		UserId = "0"
	}

	salesId, ok := c.Locals("x-key-country-id").(string)
	if !ok {
		salesId = ""
	}

	agentId, ok := c.Locals("x-key-legal-entity").(string)
	if !ok {
		agentId = ""
	}

	ctx = context.WithValue(ctx, constant.HeaderContext, model.ValueContext{
		UserId:  UserId,
		SalesId: salesId,
		AgentId: agentId,
		Token:   Token,
	})

	return context.WithValue(ctx, constant.FiberContext, c)
}

func CreateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), constant.DefaultTimeout*time.Second)
}

func CreateContextWithCustomTimeout(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}

func SetContext(ctx context.Context, c *fiber.Ctx) context.Context {
	return context.WithValue(ctx, constant.FiberContext, c)
}

func GetValueFiberFromContext(ctx context.Context) *fiber.Ctx {
	return ctx.Value(constant.FiberContext).(*fiber.Ctx)
}
