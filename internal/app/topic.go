package app

import (
	"subscription-api/internal/app/domain"
	"subscription-api/pkg/logger"
)

type Topic struct {
	service domain.Topic
	log     *logger.Logger
}

func NewTopic(service domain.Topic, log *logger.Logger) Topic {
	return Topic{
		service: service,
		log:     log,
	}
}

func (t Topic) Get() {

}

func (t Topic) Find() {

}

func (t Topic) Create() {

}

func (t Topic) Update() {

}

func (t Topic) Delete() {

}
