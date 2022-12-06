// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/miqdadyyy/jetdevs-assesment/internal/entity"
	mock "github.com/stretchr/testify/mock"

	models "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
)

// ArticleRepositoryInterface is an autogenerated mock type for the ArticleRepositoryInterface type
type ArticleRepositoryInterface struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: ctx, payload
func (_m *ArticleRepositoryInterface) CreateArticle(ctx context.Context, payload *models.Article) (*models.Article, error) {
	ret := _m.Called(ctx, payload)

	var r0 *models.Article
	if rf, ok := ret.Get(0).(func(context.Context, *models.Article) *models.Article); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Article) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticle provides a mock function with given fields: ctx, payload
func (_m *ArticleRepositoryInterface) GetArticle(ctx context.Context, payload entity.GetArticleRequest) ([]models.Article, error) {
	ret := _m.Called(ctx, payload)

	var r0 []models.Article
	if rf, ok := ret.Get(0).(func(context.Context, entity.GetArticleRequest) []models.Article); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.GetArticleRequest) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticleDetail provides a mock function with given fields: ctx, payload
func (_m *ArticleRepositoryInterface) GetArticleDetail(ctx context.Context, payload entity.GetArticleDetailRequest) (*models.Article, error) {
	ret := _m.Called(ctx, payload)

	var r0 *models.Article
	if rf, ok := ret.Get(0).(func(context.Context, entity.GetArticleDetailRequest) *models.Article); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.GetArticleDetailRequest) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalArticle provides a mock function with given fields: ctx, payload
func (_m *ArticleRepositoryInterface) GetTotalArticle(ctx context.Context, payload entity.GetArticleRequest) (int64, error) {
	ret := _m.Called(ctx, payload)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, entity.GetArticleRequest) int64); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.GetArticleRequest) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewArticleRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewArticleRepositoryInterface creates a new instance of ArticleRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArticleRepositoryInterface(t mockConstructorTestingTNewArticleRepositoryInterface) *ArticleRepositoryInterface {
	mock := &ArticleRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
