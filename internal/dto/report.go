package dto

type CreateReportRequest struct {
	CommentID int    `json:"comment_id"`
	PostID    int    `json:"post_id"`
	Reason    string `json:"reason"`
}
