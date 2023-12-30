package entity

type JobStatus int

const (
	UNKNOWN JobStatus = iota
	PENDING
	RUNNING
	FAIL
	SUCCEED
)

func (s JobStatus) String() string {
	switch s {
	case PENDING:
		return "pending"
	case RUNNING:
		return "running"
	case FAIL:
		return "fail"
	case SUCCEED:
		return "succeed"
	default:
		return "unknown"
	}
}

type IsrJob struct {
	IsrJobID                string
	UploadImageKey          string
	SuperResolutionImageKey string
}
