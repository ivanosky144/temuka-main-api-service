package publisher

import (
	"github.com/temuka-api-service/internal/constant"
	"github.com/temuka-api-service/util/queue"
)

type SearchSyncEvent struct {
	Operation string                 `json:"operation"`
	Type      string                 `json:"type"`
	EntityID  string                 `json:"entity_id"`
	Data      map[string]interface{} `json:"data"`
}

type SearchIndexPublisher interface {
	PublishSyncEvent(op, entityType, entityID string, data map[string]interface{}) error
}

type searchIndexPublisherImpl struct {
	rmq queue.RabbitMQChannel
}

func NewSearchIndexPublisher(rmq queue.RabbitMQChannel) SearchIndexPublisher {
	_ = rmq.RegisterExchange(constant.SearchExchange, "direct", true, false)
	_, _ = rmq.InitQueue(constant.SearchExchange, constant.SearcSyncRoutingKey, true, false)

	return &searchIndexPublisherImpl{rmq: rmq}
}

func (p *searchIndexPublisherImpl) PublishSyncEvent(op, entityType, entityID string, data map[string]interface{}) error {
	event := SearchSyncEvent{
		Operation: op,
		Type:      entityType,
		EntityID:  entityID,
		Data:      data,
	}

	return p.rmq.PublishMessage(constant.SearchExchange, constant.SearcSyncRoutingKey, event)
}
