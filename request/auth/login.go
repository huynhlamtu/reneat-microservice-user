package auth

type GetLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetLoginResponse struct {
	ClientUuid string `json:"client_uuid"`
	Uuid       string `json:"uuid"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Token      string `json:"token"`
}
