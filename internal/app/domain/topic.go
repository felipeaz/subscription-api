package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"subscription-api/pkg/logger"
)

type TopicServer interface {
	Get(ctx context.Context) ([]Topic, error)
	Find(ctx context.Context, id string) (Topic, error)
	Create(ctx context.Context, topicObj Topic) (Topic, error)
	Update(ctx context.Context, id string, topic Topic) error
	Delete(ctx context.Context, id string) error
	NewTopicFromRequest(c *gin.Context) (Topic, error)
}

type Topic struct {
	Name        string
	Description string
	Subscribers []*subscription
	Log         *logger.Logger
}

func (t Topic) Get(ctx context.Context) ([]Topic, error) {
	return nil, nil
}

func (t Topic) Find(ctx context.Context, id string) (Topic, error) {
	return Topic{}, nil
}

func (t Topic) Create(ctx context.Context, topicObj Topic) (Topic, error) {
	return t, nil
}

func (t Topic) Update(ctx context.Context, id string, topic Topic) error {
	return nil
}

func (t Topic) Delete(ctx context.Context, id string) error {
	return nil
}

func (t Topic) NewTopicFromRequest(c *gin.Context) (Topic, error) {
	var topicObj Topic
	err := c.BindJSON(&topicObj)
	if err != nil {
		return Topic{}, err
	}
	return topicObj, nil
}
