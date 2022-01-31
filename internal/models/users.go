package models

type User struct {
	Username string `json:"username" binding:"required,min=4,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type SignUpResponse struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}
