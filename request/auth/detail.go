package auth

type UserDetailRequest struct {
	Username string `uri:"username" binding:"required"`
}

type UserDetailResponse struct {
	ClientUuid string `json:"client_uuid"`
	Uuid       string `json:"uuid"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Name       string `json:"name"`
}
