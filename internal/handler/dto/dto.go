package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
