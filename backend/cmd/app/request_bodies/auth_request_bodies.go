package request_bodies

/* Auth Request Bodies */

type RegisterRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPwd string `json:"confirm_pwd"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

