package endpoints

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	String string `json:"message"`
	Error string `json:"error"`
}