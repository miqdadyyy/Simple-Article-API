package entity

import (
	"time"
)

type (
	GetArticleRequest struct {
		Limit    int
		Page     int
		Username string
		Search   string
	}

	GetArticleResponse struct {
		Limit    int       `json:"limit"`
		Page     int       `json:"page"`
		Total    int64     `json:"total"`
		Articles []Article `json:"articles"`
	}

	Article struct {
		ID        int64     `json:"id"`
		Username  string    `json:"username"`
		Title     string    `json:"title"`
		Slug      string    `json:"slug"`
		Thumbnail string    `json:"thumbnail"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	GetArticleDetailRequest struct {
		ID   int64
		Slug string
	}

	GetArticleDetailResponse struct {
		Article  Article   `json:"article"`
		Comments []Comment `json:"comments"`
	}

	CreateArticleRequest struct {
		Username string
		Title    string
		Content  string
	}
)
