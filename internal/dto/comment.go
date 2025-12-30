package dto

type AddCommentRequest struct {
	PostID   int    `json:"post_id"`
	UserID   int    `json:"user_id"`
	ParentID *int   `json:"parent_id"`
	Content  string `json:"content"`
}

type ShowCommentsRequest struct {
	PostID int `json:"post_id"`
}

type ShowRepliesRequest struct {
	ParentID int `json:"parent_id"`
}
