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
	Name       string `json:"name"`
	Bio        string `json:"bio"`
	ProfileUrl string `json:"profile_url"`
	Token      string `json:"token"`
}
