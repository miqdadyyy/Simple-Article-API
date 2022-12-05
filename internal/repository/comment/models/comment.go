package models

import "gorm.io/gorm"

type (
	Comment struct {
		gorm.Model
		ParentID  int64  `json:"parent_id"`
		ArticleID int64  `json:"article_id"`
		Username  string `json:"username"`
		Content   string `json:"content"`
		gorm.DeletedAt
		SubComments []Comment `json:"sub_comments" gorm:"foreignKey:parent_id;references:id"`
	}
)
