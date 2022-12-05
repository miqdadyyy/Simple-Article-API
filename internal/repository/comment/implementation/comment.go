package implementation

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	articlerm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
	commentrp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment"
	commentrm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/models"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) commentrp.CommentRepositoryInterface {
	return &CommentRepository{
		db: db,
	}
}

func (c *CommentRepository) GetAllArticleComments(
	ctx context.Context,
	articleID int64,
) ([]commentrm.Comment, error) {
	var result []commentrm.Comment
	if err := c.db.Preload("SubComments").
		Where("article_id = ?", articleID).
		Where("parent_id = 0").
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CommentRepository) CreateComment(
	ctx context.Context,
	article *articlerm.Article,
	payload entity.CreateCommentRequest,
) (*commentrm.Comment, error) {
	comment := &commentrm.Comment{
		ArticleID: int64(article.ID),
		Username:  payload.Username,
		Content:   payload.Content,
	}

	if payload.ParentID != 0 {
		var parentComment commentrm.Comment
		if err := c.db.First(&parentComment, payload.ParentID).Error; err != nil {
			return nil, err
		}

		if parentComment.ParentID != 0 {
			return nil, errors.New("Maximum nested comment is 2")
		}
		comment.ParentID = int64(parentComment.ID)
	}

	if err := c.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
