package models

type (
	UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	UserRegister struct {
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		Email      string `json:"email" binding:"required"`
		ValidUntil string `json:"valid_until" binding:"required"`
		FullName   string `json:"full_name" binding:"required"`
	}
)
