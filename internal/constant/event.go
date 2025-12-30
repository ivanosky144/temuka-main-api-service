package constant

const (
	SearchExchange         = "temuka_search_exchange"
	RecommendationExchange = "temuka_recommendation_exchange"
	AnalyticsExchange      = "temuka_analytics_exchange"

	SearcSyncRoutingKey            = "search.sync"
	RecommendationUpdateRoutingKey = "recommendation.update"
	AnalyticsEventRoutingKey       = "analytics.event"

	EventOperationCreate = "CREATE"
	EventOperationDelete = "DELETE"
	EventOperationUpdate = "UPDATE"

	EventEntityTypePost       = "POST"
	EventEntityTypeUser       = "USER"
	EventEntityTypeCommunity  = "COMMUNITY"
	EventEntityTypeUniversity = "UNIVERSITY"
	EventEntityTypeMajor      = "MAJOR"
)
