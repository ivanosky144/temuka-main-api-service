package constant

const (
	ExchangeName = "temuka_exchange"

	PostCreatedRoutingKey = "post.created"
	PostUpdatedRoutingKey = "post.updated"
	PostDeletedRoutingKey = "post.deleted"
	PostLikedRoutingKey   = "post.liked"
	PostViewedRoutingKey  = "post.viewed"

	UserCreatedRoutingKey    = "user.created"
	UserUpdatedRoutingKey    = "user.updated"
	UserFollowedRoutingKey   = "user.followed"
	UserUnfollowedRoutingKey = "user.unfollowed"

	CommentCreatedRoutingKey = "comment.created"
	CommentDeletedRoutingKey = "comment.deleted"

	CommunityViewedRoutingKey = "community.viewed"
	CommunityJoinedRoutingKey = "community.joined"
	CommunityLeftRoutingKey   = "community.left"

	UniversityViewedRoutingKey   = "university.viewed"
	UniversityReviewedRoutingKey = "university.reviewed"

	MajorViewedRoutingKey   = "major.viewed"
	MajorReviewedRoutingKey = "major.reviewed"
)
