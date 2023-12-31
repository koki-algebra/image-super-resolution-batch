package entity

import "time"

type History struct {
	HistoryID int `bun:",pk"`
	Timestamp time.Time
	Status    JobStatus
	IsrJobID  string
	IsrJob    IsrJob `bun:"rel:belongs-to"`
}
