package model

import "time"

type Project struct {
	Id        string
	UserId    string
	Name      string
	CreatedAt time.Time
}
