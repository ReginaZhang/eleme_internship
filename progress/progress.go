package progress

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/utils"
	"golang.org/x/net/websocket"
)

// Progress keeps track of progress of all jobs
type Progress interface {
	// AddNewJob creates a new JobProgress with jobID if not exists
	AddNewJob(jobID string, numOfRuns int) error

	// Add handles a callback event. It will return true if the event marks
	// the end of a job.
	Add(utils.Map) (done bool, err error)

	// Stream sends all events to the given websocket
	Stream(jobID string, ws *websocket.Conn) error

	// Summaries returns summaries of the given job
	Summaries(jobID string) (summaries map[*models.Statistic][]*models.Failure, failed bool, err error)

	// Clean remove progress of the given job
	Clean(jobID string)
}
