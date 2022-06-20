package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"subscription-api/internal/app/domain"
)

type Topic struct {
	service domain.TopicServer
}

func NewTopic(service domain.TopicServer) Topic {
	return Topic{
		service: service,
	}
}

func (t Topic) Get(c *gin.Context) {
	ctx := context.Background()
	topics, err := t.service.Get(ctx)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (t Topic) Find(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	topic, err := t.service.Find(ctx, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, topic)
}

func (t Topic) Create(c *gin.Context) {
	ctx := context.Background()
	topic, err := t.service.NewTopicFromRequest(c)
	if err != nil {
		return
	}
	topicCreated, err := t.service.Create(ctx, topic)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, topicCreated)
}

func (t Topic) Update(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	topic, err := t.service.NewTopicFromRequest(c)
	if err != nil {
		return
	}
	err = t.service.Update(ctx, id, topic)
	if err != nil {
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (t Topic) Delete(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	err := t.service.Delete(ctx, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
