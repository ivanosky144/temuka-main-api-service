package dto

type SearchUsersDTO struct {
	Name string `json:"name"`
}

type GetUserDetailDTO struct {
	UserID int `json:"user_id"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	UserID         int    `json:"user_id"`
	Username       string `json:"username"`
	Desc           string `json:"desc"`
	Displayname    string `json:"displayname"`
	ProfilePicture string `json:"profile_picture"`
}

type FollowUserDTO struct {
	TargetID      int `json:"target_id"`
	CurrentUserID int `json:"currentuser_id"`
}

type GetFollowersDTO struct {
	UserID int `json:"user_id"`
}
