package auth

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
	ClientUuid string `json:"client_uuid"`
	Uuid       string `json:"uuid"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}
