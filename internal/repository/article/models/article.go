package models

import (
	"gorm.io/gorm"

	commentrpm "github.com/miqdadyyy/jetdevs-assesment/internal/repository/comment/models"
)

type (
	Article struct {
		gorm.Model
		Username  string `json:"username"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		Content   string `json:"content"`
		Thumbnail string `json:"thumbnail"`
		gorm.DeletedAt
		Comments []commentrpm.Comment `json:"comments"`
	}
)
