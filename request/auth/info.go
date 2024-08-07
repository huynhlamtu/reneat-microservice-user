package auth

type UserInfoRequest struct {
	Uuid string `uri:"uuid" binding:"required"`
}

type UserInfoResponse struct {
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
