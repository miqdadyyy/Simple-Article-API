package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"
	"gorm.io/gorm"

	"github.com/miqdadyyy/jetdevs-assesment/internal/constant"
	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
)

func (h *Handler) HandleGetArticles(ctx *fiber.Ctx) error {
	var payload entity.GetArticleRequest
	if err := ctx.QueryParser(&payload); err != nil {
		return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
	}

	data, err := h.articleUsecase.GetArticlesByPayload(ctx.Context(), payload)
	if err != nil {
		logger.Errorf("failed to get articles with err : %v", err)
		return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
	}

	return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusOK, data))
}

func (h *Handler) HandleCreateArticle(ctx *fiber.Ctx) error {
	var payload entity.CreateArticleRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
	}

	data, err := h.articleUsecase.CreateArticle(ctx.Context(), payload)
	if err != nil {
		logger.Errorf("failed to create articles with err : %v", err)
		return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
	}

	return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusOK, data))
}

func (h *Handler) HandleGetArticleDetail(ctx *fiber.Ctx) error {
	var payload entity.GetArticleDetailRequest
	if err := ctx.ParamsParser(&payload); err != nil {
		return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
	}

	data, err := h.articleUsecase.GetArticleDetail(ctx.Context(), payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusNotFound, nil))
		} else {
			return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
		}
	}

	return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusOK, data))
}

func (h *Handler) HandleCreateArticleComment(ctx *fiber.Ctx) error {
	var payload entity.CreateCommentRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
	}
	payload.ArticleSlug = ctx.Params("slug")

	if err := h.articleUsecase.CreateComment(ctx.Context(), payload); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusNotFound, nil))
		} else {
			return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusBadRequest, nil))
		}
	}

	return ctx.JSON(entity.NewHTTPJsonResponse(constant.HTTPStatusOK, nil))
}
