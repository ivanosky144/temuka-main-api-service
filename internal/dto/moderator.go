package dto

type SendModeratorRequest struct {
	CommunityID       int `json:"community_id"`
	CommunityMemberID int `json:"communitymember_id"`
}
