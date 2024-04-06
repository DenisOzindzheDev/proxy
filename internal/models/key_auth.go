package models

import "time"

type KeyAuth struct {
	id        string
	createdAt time.Time
	Consumer  *Consumer
	Service   *Service
}

func NewKeyAuth(consumer *Consumer, service *Service) (*KeyAuth, error) {
	return &KeyAuth{Consumer: consumer, Service: service}, nil
}

// keyauth pair
func CheckAuth() bool {
	return true
}
