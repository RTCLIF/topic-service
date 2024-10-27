package model

import "time"

type TopicInfo struct {
	Title       string
	Description string
}

type Topic struct {
	Id        string
	Info      TopicInfo
	Status    string
	CreatedAt time.Time
	UpdatedAt *time.Time // если изменений ещё не было то nil
}
