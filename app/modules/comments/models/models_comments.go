package models

type (
	UpdateComment struct {
		CommentBody string `json:"comments_body"`
	}
	GetAllComment struct {
		CommentBody string `json:"comments_body"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
	}
)
