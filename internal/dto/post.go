package dto

type CreatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
	CommunityID int    `json:"community_id"`
}

type UpdatePostRequest struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type LikePostRequest struct {
	UserID int `json:"user_id"`
}

type PostCreatedEventData struct {
	PostID      int    `json:"post_id"`
	UserID      int    `json:"user_id"`
	CommunityID int    `json:"community_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostLikedEventData struct {
	PostID        int `json:"post_id"`
	PostOwnerID   int `json:"post_owner_id"`
	LikedByUserID int `json:"liked_by_user_id"`
	CommunityID   int `json:"community_id"`
}
