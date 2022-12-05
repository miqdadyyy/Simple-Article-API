package comment

import (
	"context"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	articlerm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
	commentrm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/models"
)

type CommentRepositoryInterface interface {
	GetAllArticleComments(ctx context.Context, articleID int64) ([]commentrm.Comment, error)
	CreateComment(ctx context.Context, article *articlerm.Article, payload entity.CreateCommentRequest) (*commentrm.Comment, error)
}
