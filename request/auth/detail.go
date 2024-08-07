package auth

type UserDetailRequest struct {
	Username string `uri:"username" binding:"required"`
}

type UserDetailResponse struct {
	ClientUuid string   `json:"client_uuid"`
	Uuid       string   `json:"uuid"`
	Email      string   `json:"email"`
	Username   string   `json:"username"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Name       string   `json:"name"`
	ProfileUrl string   `json:"profile_url"`
	Bio        string   `json:"bio"`
	Followers  []string `json:"followers"`
	Followings []string `json:"followings"`
	Posts      []string `json:"posts"`
}
