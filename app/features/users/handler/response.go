package handler

type LoginResponse struct {
	Phone string `json:"phone"`
	Nama  string `json:"nama"`
	Token string `json:"token"`
}

type UpdateResponse struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
