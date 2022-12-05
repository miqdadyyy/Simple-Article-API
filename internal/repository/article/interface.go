package article

import (
	"context"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	"github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
)

type ArticleRepositoryInterface interface {
	GetArticle(ctx context.Context, payload entity.GetArticleRequest) ([]models.Article, error)
	GetTotalArticle(ctx context.Context, payload entity.GetArticleRequest) (int64, error)
	GetArticleDetail(ctx context.Context, payload entity.GetArticleDetailRequest) (*models.Article, error)
	CreateArticle(ctx context.Context, payload *models.Article) (*models.Article, error)
}
