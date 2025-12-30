package dto

type CreateCommunityRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LogoPicture  string `json:"logo_picture"`
	CoverPicture string `json:"cover_picture"`
}

type UpdateCommunityRequest struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	LogoPicture  string `json:"logo_picture"`
	CoverPicture string `json:"cover_picture"`
}

type JoinCommunityRequest struct {
	UserID int `json:"user_id"`
}

type GetUserJoinedCommunitiesRequest struct {
	UserID int `json:"user_id"`
}
