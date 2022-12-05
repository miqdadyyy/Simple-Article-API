package entity

import "time"

type (
	Comment struct {
		ID          int64     `json:"id"`
		SubComments []Comment `json:"sub_comments"`
		Username    string    `json:"username"`
		Content     string    `json:"content"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	CreateCommentRequest struct {
		ArticleSlug string `json:"article_id"`
		ParentID    int64  `json:"parent_id"`
		Username    string `json:"username"`
		Content     string `json:"content"`
	}
)
