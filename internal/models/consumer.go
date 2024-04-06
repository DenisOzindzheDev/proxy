package models

import (
	"rproxy/pkg/util"
	"time"
)

type Consumer struct {
	id        string
	createdAt time.Time
	Name      string
	Services []Service
}

func NewConsumer(name string) *Consumer {
	return &Consumer{id: util.GenerateID(), createdAt: time.Unix(time.Now().Unix(), 0), Name: name, Services: make([]Service, 0}
}
