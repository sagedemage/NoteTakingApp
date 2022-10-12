package request_bodies

/* Registration */

type RegisterRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPwd string `json:"confirm_pwd"`
}

/* Login */

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

