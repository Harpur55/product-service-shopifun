package handler

import (
	"codebase-app/internal/adapter"
	"codebase-app/internal/middleware"
	"codebase-app/internal/module/product/entity"
	ports "codebase-app/internal/module/product/port"
	"codebase-app/internal/module/product/repository"
	"codebase-app/internal/module/product/service"
	"codebase-app/pkg/errmsg"
	"codebase-app/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type productHandler struct {
	service ports.ProductService
}

func NewProductHandler() *productHandler {
	var (
		handler = new(productHandler)
		repo    = repository.NewProductRepository(adapter.Adapters.ShopeefunPostgres)
		service = service.NewProductService(repo)
	)
	handler.service = service

	return handler
}

func (h *productHandler) Register(router fiber.Router) {
	//  router.Get("/shops", middleware.UserIdHeader, h.GetShops)
	router.Post("/product", middleware.ProductIdHeader, h.CreateProduct)
	router.Get("/product", h.GetProduct)
	router.Get("/product/:id", h.GetProductByid)
	 router.Patch("/product/:id", middleware.ProductIdHeader, h.UpdateProduct)
	router.Delete("/product/:id", middleware.ProductIdHeader, h.DeleteProduct)
	
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.ProductId = l.ProductId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::CreateProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))

}

func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		req = new(entity.GetProductRequest)
	)

	resp, err := h.service.GetProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *productHandler) GetProductByid(c *fiber.Ctx) error {
	var (
		req = new(entity.GetProductByidRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
	)

	req.ProductId = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::GetProductbyid - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProductById(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.DeleteProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		 l   = middleware.GetLocals(c)
	)
	 req.ProductId = l.ProductId
	req.ProductId = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::DeleteProduct - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(nil, ""))
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateProductRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateProduct - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.ProductId = l.ProductId
	req.ProductId = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::UpdateShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.UpdateProduct(ctx, req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))
}

// func (h *shopHandler) GetShops(c *fiber.Ctx) error {
// 	var (
// 		req = new(entity.ShopsRequest)
// 		ctx = c.Context()
// 		v   = adapter.Adapters.Validator
// 		l   = middleware.GetLocals(c)
// 	)

// 	if err := c.QueryParser(req); err != nil {
// 		log.Warn().Err(err).Msg("handler::GetShops - Parse request query")
// 		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
// 	}

// 	req.UserId = l.UserId
// 	req.SetDefault()

// 	if err := v.Validate(req); err != nil {
// 		log.Warn().Err(err).Any("payload", req).Msg("handler::GetShops - Validate request body")
// 		code, errs := errmsg.Errors(err, req)
// 		return c.Status(code).JSON(response.Error(errs))
// 	}

// 	resp, err := h.service.GetShops(ctx, req)
// 	if err != nil {
// 		code, errs := errmsg.Errors[error](err)
// 		return c.Status(code).JSON(response.Error(errs))
// 	}

// 	return c.Status(fiber.StatusOK).JSON(response.Success(resp, ""))

// }
