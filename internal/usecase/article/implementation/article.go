package implementation

import (
	"context"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	articlerp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article"
	articlerpm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
	commentrp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment"
	articleuc "github.com/miqdadyyy/jetdevs-assesment/internal/usecase/article"
	"github.com/miqdadyyy/jetdevs-assesment/pkg/str"
)

type ArticleUsecase struct {
	articleRepository articlerp.ArticleRepositoryInterface
	commentRepository commentrp.CommentRepositoryInterface
}

func NewArticleUsecase(
	articleRepository articlerp.ArticleRepositoryInterface,
	commentRepository commentrp.CommentRepositoryInterface,
) articleuc.ArticleUsecaseInterface {
	return &ArticleUsecase{
		articleRepository: articleRepository,
		commentRepository: commentRepository,
	}
}

func (a *ArticleUsecase) GetArticlesByPayload(
	ctx context.Context,
	payload entity.GetArticleRequest,
) (*entity.GetArticleResponse, error) {
	var articles []entity.Article
	if payload.Limit <= 0 {
		payload.Limit = 20
	}

	if payload.Page <= 0 {
		payload.Page = 1
	}

	articleData, err := a.articleRepository.GetArticle(ctx, payload)
	if err != nil {
		return nil, err
	}

	for _, article := range articleData {
		articles = append(articles, entity.Article{
			ID:        int64(article.ID),
			Username:  article.Username,
			Title:     article.Title,
			Slug:      article.Slug,
			Content:   str.StringEllipsis(article.Content, 100),
			Thumbnail: article.Thumbnail,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		})
	}

	total, err := a.articleRepository.GetTotalArticle(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &entity.GetArticleResponse{
		Limit:    payload.Limit,
		Page:     payload.Page,
		Total:    total,
		Articles: articles,
	}, nil
}

func (a *ArticleUsecase) CreateArticle(
	ctx context.Context,
	payload entity.CreateArticleRequest,
) (*articlerpm.Article, error) {
	article, err := a.articleRepository.CreateArticle(ctx, &articlerpm.Article{
		Username: payload.Username,
		Title:    payload.Title,
		Content:  payload.Content,
	})
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleUsecase) GetArticleDetail(
	ctx context.Context,
	payload entity.GetArticleDetailRequest,
) (*entity.GetArticleDetailResponse, error) {
	var article entity.Article
	comments := make([]entity.Comment, 0)
	articleData, err := a.articleRepository.GetArticleDetail(ctx, payload)
	if err != nil {
		return nil, err
	}
	article = entity.Article{
		ID:        int64(articleData.ID),
		Username:  articleData.Username,
		Title:     articleData.Title,
		Slug:      articleData.Slug,
		Content:   articleData.Content,
		Thumbnail: articleData.Thumbnail,
		CreatedAt: articleData.CreatedAt,
		UpdatedAt: articleData.UpdatedAt,
	}

	commentRows, err := a.commentRepository.GetAllArticleComments(ctx, article.ID)
	if err != nil {
		return nil, err
	}
	for _, commentRow := range commentRows {
		subComments := make([]entity.Comment, 0)
		for _, subComment := range commentRow.SubComments {
			subComments = append(subComments, entity.Comment{
				ID:        int64(subComment.ID),
				Username:  subComment.Username,
				Content:   subComment.Content,
				CreatedAt: subComment.CreatedAt,
				UpdatedAt: subComment.UpdatedAt,
			})
		}

		comments = append(comments, entity.Comment{
			ID:          int64(commentRow.ID),
			Username:    commentRow.Username,
			Content:     commentRow.Content,
			CreatedAt:   commentRow.CreatedAt,
			UpdatedAt:   commentRow.UpdatedAt,
			SubComments: subComments,
		})
	}

	return &entity.GetArticleDetailResponse{
		Article:  article,
		Comments: comments,
	}, nil
}

func (a *ArticleUsecase) CreateComment(ctx context.Context, payload entity.CreateCommentRequest) error {
	article, err := a.articleRepository.GetArticleDetail(ctx, entity.GetArticleDetailRequest{Slug: payload.ArticleSlug})
	if err != nil {
		return err
	}

	_, err = a.commentRepository.CreateComment(ctx, article, payload)
	return err
}
