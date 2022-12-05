package implementation

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	"github.com/miqdadyyy/jetdevs-assesment/internal/repository/article"
	"github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
	"github.com/miqdadyyy/jetdevs-assesment/pkg/str"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) article.ArticleRepositoryInterface {
	return &ArticleRepository{
		db: db,
	}
}

func (a *ArticleRepository) GetArticle(
	ctx context.Context,
	payload entity.GetArticleRequest,
) ([]models.Article, error) {
	var result []models.Article
	query := a.db.
		Limit(payload.Limit).
		Offset(payload.Limit * (payload.Page - 1)).
		Order("id desc")

	if payload.Search != "" {
		query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", payload.Search))
	}

	if err := query.
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ArticleRepository) GetTotalArticle(
	ctx context.Context,
	payload entity.GetArticleRequest,
) (int64, error) {
	var result int64
	if err := a.db.Model(&models.Article{}).
		Count(&result).
		Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (a *ArticleRepository) GetArticleDetail(
	ctx context.Context,
	payload entity.GetArticleDetailRequest,
) (*models.Article, error) {
	var result *models.Article
	query := a.db
	if payload.ID != 0 {
		query = query.Where("id = ?", payload.ID)
	}

	if payload.Slug != "" {
		query = query.Where("slug = ?", payload.Slug)
	}

	if err := query.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *ArticleRepository) CreateArticle(
	ctx context.Context,
	payload *models.Article,
) (*models.Article, error) {
	payload.Slug = fmt.Sprintf("%s-%s", slug.Make(payload.Title), uuid.NewString()[:8])
	payload.Thumbnail = str.GenerateRandomAvatar(payload.Slug)
	if err := a.db.Save(&payload).Error; err != nil {
		return nil, err
	}

	return payload, nil
}
