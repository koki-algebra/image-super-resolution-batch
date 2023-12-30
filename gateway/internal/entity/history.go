package entity

import "time"

type History struct {
	HistoryID int
	Timestamp time.Time
	Status    JobStatus
	IsrJobID  int
	IsrJob    IsrJob
}
