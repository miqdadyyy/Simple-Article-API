package implementation

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	articlerp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article"
	articlerpmock "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/mocks"
	articlerpm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
	commentrp "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment"
	commentrpmock "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/mocks"
	commentrpm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/models"
	"github.com/miqdadyyy/jetdevs-assesment/internal/usecase/article"
)

var mockArticleRepository = &articlerpmock.ArticleRepositoryInterface{}
var mockCommentRepository = &commentrpmock.CommentRepositoryInterface{}
var mockCtx = context.Background()
var mockError = errors.New("dummy error")

func TestArticleUsecase_CreateArticle(t *testing.T) {
	type fields struct {
		articleRepository articlerp.ArticleRepositoryInterface
		commentRepository commentrp.CommentRepositoryInterface
	}
	type args struct {
		ctx     context.Context
		payload entity.CreateArticleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *articlerpm.Article
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "failed create article",
			fields: fields{
				articleRepository: mockArticleRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.CreateArticleRequest{
					Username: "Test",
					Title:    "Test",
					Content:  "Lorem Ipsum",
				},
			},
			wantErr: true,
			want:    nil,
			mock: func(args args) {
				mockArticleRepository.On("CreateArticle", args.ctx, &articlerpm.Article{
					Username: args.payload.Username,
					Title:    args.payload.Title,
					Content:  args.payload.Content,
				}).
					Return(&articlerpm.Article{}, mockError).Once()
			},
		},
		{
			name: "success create article",
			fields: fields{
				articleRepository: mockArticleRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.CreateArticleRequest{
					Username: "Test",
					Title:    "Test",
					Content:  "Lorem Ipsum",
				},
			},
			wantErr: false,
			want: &articlerpm.Article{
				Username: "Test",
				Title:    "Test",
				Content:  "Lorem Ipsum",
			},
			mock: func(args args) {
				mockArticleRepository.On("CreateArticle", args.ctx, &articlerpm.Article{
					Username: args.payload.Username,
					Title:    args.payload.Title,
					Content:  args.payload.Content,
				}).Return(&articlerpm.Article{
					Username: "Test",
					Title:    "Test",
					Content:  "Lorem Ipsum",
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleUsecase{
				articleRepository: tt.fields.articleRepository,
				commentRepository: tt.fields.commentRepository,
			}
			tt.mock(tt.args)
			got, err := a.CreateArticle(tt.args.ctx, tt.args.payload)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Username, got.Username)
				assert.Equal(t, tt.want.Content, got.Content)
			}
		})
	}
}

func TestArticleUsecase_CreateComment(t *testing.T) {
	type fields struct {
		articleRepository articlerp.ArticleRepositoryInterface
		commentRepository commentrp.CommentRepositoryInterface
	}
	type args struct {
		ctx     context.Context
		payload entity.CreateCommentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "success create comment",
			fields: fields{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.CreateCommentRequest{
					ArticleSlug: "test-slug",
					Username:    "test",
					Content:     "test content",
				},
			},
			wantErr: false,
			mock: func(args args) {
				mockArticleRepository.On(
					"GetArticleDetail",
					mockCtx,
					entity.GetArticleDetailRequest{Slug: args.payload.ArticleSlug},
				).Return(&articlerpm.Article{
					Model: gorm.Model{ID: 1},
				}, nil).Once()

				mockCommentRepository.On(
					"CreateComment",
					mockCtx,
					&articlerpm.Article{
						Model: gorm.Model{ID: 1},
					},
					args.payload,
				).Return(&commentrpm.Comment{}, nil).Once()
			},
		},
		{
			name: "failed get article",
			fields: fields{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.CreateCommentRequest{
					ArticleSlug: "test-slug",
					Username:    "test",
					Content:     "test content",
				},
			},
			wantErr: true,
			mock: func(args args) {
				mockArticleRepository.On(
					"GetArticleDetail",
					mockCtx,
					entity.GetArticleDetailRequest{Slug: args.payload.ArticleSlug},
				).Return(&articlerpm.Article{}, mockError).Once()
			},
		},
		{
			name: "failed create comment",
			fields: fields{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.CreateCommentRequest{
					ArticleSlug: "test-slug",
					Username:    "test",
					Content:     "test content",
				},
			},
			wantErr: false,
			mock: func(args args) {
				mockArticleRepository.On(
					"GetArticleDetail",
					mockCtx,
					entity.GetArticleDetailRequest{Slug: args.payload.ArticleSlug},
				).Return(&articlerpm.Article{
					Model: gorm.Model{ID: 1},
				}, nil).Once()

				mockCommentRepository.On(
					"CreateComment",
					mockCtx,
					&articlerpm.Article{
						Model: gorm.Model{ID: 1},
					},
					args.payload,
				).Return(&commentrpm.Comment{}, mockError).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleUsecase{
				articleRepository: tt.fields.articleRepository,
				commentRepository: tt.fields.commentRepository,
			}
			tt.mock(tt.args)
			err := a.CreateComment(tt.args.ctx, tt.args.payload)
			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestArticleUsecase_GetArticleDetail(t *testing.T) {
	type fields struct {
		articleRepository articlerp.ArticleRepositoryInterface
		commentRepository commentrp.CommentRepositoryInterface
	}
	type args struct {
		ctx     context.Context
		payload entity.GetArticleDetailRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.GetArticleDetailResponse
		wantErr bool
		mock    func(args args)
	}{
		{
			name: "success get article detail",
			fields: fields{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.GetArticleDetailRequest{
					Slug: "test-slug",
				},
			},
			wantErr: false,
			want: &entity.GetArticleDetailResponse{
				Article: entity.Article{
					ID: 1,
				},
				Comments: []entity.Comment{
					{
						Username: "test",
						Content:  "test",
					},
				},
			},
			mock: func(args args) {
				mockArticleRepository.On("GetArticleDetail", mockCtx, args.payload).
					Return(&articlerpm.Article{
						Model: gorm.Model{ID: 1},
					}, nil).Once()

				mockCommentRepository.On("GetAllArticleComments", mockCtx, int64(1)).Return([]commentrpm.Comment{
					{
						Username: "test",
						Content:  "test",
						SubComments: []commentrpm.Comment{
							{
								Username: "test-sub",
								Content:  "test-sub",
							},
						},
					},
				}, nil).Once()
			},
		},
		{
			name: "failed get article detail",
			fields: fields{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.GetArticleDetailRequest{
					Slug: "test-slug",
				},
			},
			wantErr: true,
			mock: func(args args) {
				mockArticleRepository.On("GetArticleDetail", mockCtx, args.payload).
					Return(&articlerpm.Article{
						Model: gorm.Model{ID: 1},
					}, mockError).Once()
			},
		},
		{
			name: "success get article comments",
			fields: fields{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			args: args{
				ctx: mockCtx,
				payload: entity.GetArticleDetailRequest{
					Slug: "test-slug",
				},
			},
			wantErr: true,
			mock: func(args args) {
				mockArticleRepository.On("GetArticleDetail", mockCtx, args.payload).
					Return(&articlerpm.Article{
						Model: gorm.Model{ID: 1},
					}, nil).Once()

				mockCommentRepository.On("GetAllArticleComments", mockCtx, int64(1)).Return([]commentrpm.Comment{
					{
						Username: "test",
						Content:  "test",
					},
				}, mockError).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleUsecase{
				articleRepository: tt.fields.articleRepository,
				commentRepository: tt.fields.commentRepository,
			}
			tt.mock(tt.args)
			got, err := a.GetArticleDetail(tt.args.ctx, tt.args.payload)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want.Article.ID, got.Article.ID)
				assert.Equal(t, len(tt.want.Comments), len(got.Comments))
			}
		})
	}
}

func TestArticleUsecase_GetArticlesByPayload(t *testing.T) {
	type fields struct {
		articleRepository articlerp.ArticleRepositoryInterface
		commentRepository commentrp.CommentRepositoryInterface
	}
	type args struct {
		ctx     context.Context
		payload entity.GetArticleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.GetArticleResponse
		wantErr bool
		mock    func(args args)
	}{
		{
			name:   "success get articles",
			fields: fields{articleRepository: mockArticleRepository},
			args: args{
				ctx:     mockCtx,
				payload: entity.GetArticleRequest{},
			},
			want: &entity.GetArticleResponse{
				Limit: 20,
				Page:  1,
				Total: 1,
				Articles: []entity.Article{
					{
						ID: 1,
					},
				},
			},
			wantErr: false,
			mock: func(args args) {
				mockArticleRepository.On("GetArticle", mockCtx, entity.GetArticleRequest{
					Limit: 20,
					Page:  1,
				}).Return([]articlerpm.Article{
					{
						Model: gorm.Model{ID: 1},
					},
				}, nil).Once()

				mockArticleRepository.On("GetTotalArticle", mockCtx, entity.GetArticleRequest{
					Limit: 20,
					Page:  1,
				}).Return(int64(1), nil).Once()
			},
		},
		{
			name:   "failed get articles",
			fields: fields{articleRepository: mockArticleRepository},
			args: args{
				ctx: mockCtx,
				payload: entity.GetArticleRequest{
					Limit: 20,
					Page:  1,
				},
			},
			wantErr: true,
			mock: func(args args) {
				mockArticleRepository.On("GetArticle", mockCtx, args.payload).
					Return([]articlerpm.Article{
						{
							Model: gorm.Model{ID: 1},
						},
					}, mockError).Once()
			},
		},
		{
			name:   "success get total articles",
			fields: fields{articleRepository: mockArticleRepository},
			args: args{
				ctx: mockCtx,
				payload: entity.GetArticleRequest{
					Limit: 20,
					Page:  1,
				},
			},
			wantErr: true,
			mock: func(args args) {
				mockArticleRepository.On("GetArticle", mockCtx, args.payload).
					Return([]articlerpm.Article{
						{
							Model: gorm.Model{ID: 1},
						},
					}, nil).Once()

				mockArticleRepository.On("GetTotalArticle", mockCtx, args.payload).
					Return(int64(1), mockError).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArticleUsecase{
				articleRepository: tt.fields.articleRepository,
				commentRepository: tt.fields.commentRepository,
			}
			tt.mock(tt.args)
			got, err := a.GetArticlesByPayload(tt.args.ctx, tt.args.payload)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want.Page, got.Page)
				assert.Equal(t, tt.want.Limit, got.Limit)
				assert.Equal(t, tt.want.Total, got.Total)
			}
		})
	}
}

func TestNewArticleUsecase(t *testing.T) {
	type args struct {
		articleRepository articlerp.ArticleRepositoryInterface
		commentRepository commentrp.CommentRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want article.ArticleUsecaseInterface
	}{
		{
			name: "success",
			args: args{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
			want: &ArticleUsecase{
				articleRepository: mockArticleRepository,
				commentRepository: mockCommentRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArticleUsecase(tt.args.articleRepository, tt.args.commentRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
