package article

import (
	"context"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	articlerpm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
)

type ArticleUsecaseInterface interface {
	GetArticlesByPayload(ctx context.Context, payload entity.GetArticleRequest) (*entity.GetArticleResponse, error)
	GetArticleDetail(ctx context.Context, payload entity.GetArticleDetailRequest) (*entity.GetArticleDetailResponse, error)
	CreateArticle(ctx context.Context, payload entity.CreateArticleRequest) (*articlerpm.Article, error)
	CreateComment(ctx context.Context, payload entity.CreateCommentRequest) error
}
