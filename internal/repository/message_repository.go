package repository

import (
	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/couchbase/gocb/v2"
)

type MessageRepository struct {
	Cluster *gocb.Cluster
}

func NewMessageRepository(cluster *gocb.Cluster) *MessageRepository {
	return &MessageRepository{
		Cluster: cluster,
	}
}

const MessageLimit = 2

func (r *MessageRepository) GetMessagesByStatus(status model.Status, limit int) ([]model.Message, error) {
	query := "SELECT u.* FROM messages AS u WHERE u.status = $status LIMIT $limit"
	result, err := r.Cluster.Query(query, &gocb.QueryOptions{
		NamedParameters: map[string]any{
			"status": status,
			"limit":  limit,
		},
	})
	if err != nil {
		return nil, err
	}

	var messages []model.Message
	for result.Next() {
		var msg model.Message
		if err := result.Row(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) UpdateMessageStatus(msg model.Message, status model.Status) (model.Message, error) {
	msg.Status = status

	_, err := r.Cluster.Bucket("messages").DefaultCollection().Upsert(msg.ID, msg, nil)
	if err != nil {
		return msg, err
	}

	return msg, nil
}
