package models

import "time"

type Service struct {
	uuid      string
	createdAt time.Time
	Name      string
	Routes    []Routes
}
