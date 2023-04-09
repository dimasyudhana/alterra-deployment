package handlers

type RegisterInput struct {
	Username         string `json:"username"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirm_password"`
}

type LoginInput struct {
	Phone            string `json:"phone"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirm_password"`
}

type UpdateUsername struct {
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
