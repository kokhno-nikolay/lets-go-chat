package models

type User struct {
	Username string `json:"username" binding:"required,min=4,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	UUID     string `json:"uuid,omitempty"`
}

type SignUpResponse struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

type SignInResponse struct {
	URL string `json:"url"`
}
