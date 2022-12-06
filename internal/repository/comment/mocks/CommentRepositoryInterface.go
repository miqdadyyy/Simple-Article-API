// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	commentmodels "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/models"

	entity "github.com/miqdadyyy/jetdevs-assesment/internal/entity"

	mock "github.com/stretchr/testify/mock"

	models "github.com/miqdadyyy/jetdevs-assesment/internal/repository/article/models"
)

// CommentRepositoryInterface is an autogenerated mock type for the CommentRepositoryInterface type
type CommentRepositoryInterface struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: ctx, article, payload
func (_m *CommentRepositoryInterface) CreateComment(ctx context.Context, article *models.Article, payload entity.CreateCommentRequest) (*commentmodels.Comment, error) {
	ret := _m.Called(ctx, article, payload)

	var r0 *commentmodels.Comment
	if rf, ok := ret.Get(0).(func(context.Context, *models.Article, entity.CreateCommentRequest) *commentmodels.Comment); ok {
		r0 = rf(ctx, article, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commentmodels.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Article, entity.CreateCommentRequest) error); ok {
		r1 = rf(ctx, article, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllArticleComments provides a mock function with given fields: ctx, articleID
func (_m *CommentRepositoryInterface) GetAllArticleComments(ctx context.Context, articleID int64) ([]commentmodels.Comment, error) {
	ret := _m.Called(ctx, articleID)

	var r0 []commentmodels.Comment
	if rf, ok := ret.Get(0).(func(context.Context, int64) []commentmodels.Comment); ok {
		r0 = rf(ctx, articleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]commentmodels.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentRepositoryInterface creates a new instance of CommentRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentRepositoryInterface(t mockConstructorTestingTNewCommentRepositoryInterface) *CommentRepositoryInterface {
	mock := &CommentRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
