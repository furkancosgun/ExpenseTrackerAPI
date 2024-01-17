package model

import "time"

type Project struct {
	ProjectId string
	UserId    string
	Name      string
	CreatedAt time.Time
}
