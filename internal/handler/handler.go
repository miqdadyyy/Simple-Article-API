package handler

import "github.com/miqdadyyy/jetdevs-assesment/internal/usecase/article"

type Handler struct {
	articleUsecase article.ArticleUsecaseInterface
}

func NewHandler(
	articleUsecase article.ArticleUsecaseInterface,
) *Handler {
	return &Handler{
		articleUsecase: articleUsecase,
	}
}
