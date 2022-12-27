package models

type Registration struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
