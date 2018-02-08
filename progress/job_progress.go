package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/Sirupsen/logrus"
	"github.com/hpcloud/tail"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

type jobProgress struct {
	sync.RWMutex
	ID        string
	file      *os.File
	NumOfRuns int
	Runs      map[string]*RunResult // map[RunID]*RunResult
	done      chan bool
	StartTime time.Time
	failed    bool
}

func newJobProgress(jobID string, startTime time.Time, numOfRuns int) (*jobProgress, error) {
	f, err := os.OpenFile(fmt.Sprintf("output_%s.tmp", jobID), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open temporary output file")
	}
	return &jobProgress{
		ID:        jobID,
		file:      f,
		NumOfRuns: numOfRuns,
		Runs:      make(map[string]*RunResult),
		done:      make(chan bool),
		StartTime: startTime,
	}, nil
}

func (j *jobProgress) Write(b []byte) error {
	j.Lock()
	_, err := j.file.Write(b)
	j.Unlock()
	return err
}

func (j *jobProgress) Add(body utils.Map) (done bool, err error) {
	c, err := json.Marshal(body)
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal event")
	}
	c = append(c, byte('\n'))
	if err := j.Write(c); err != nil {
		return false, errors.Wrap(err, "failed to write to tmp file")
	}

	switch (body)["type"] {
	case "failed":
		if err := j.RecordFailure(body); err != nil {
			return false, errors.Wrap(err, "failed to record failure")
		}
	case "stats":
		j.RecordStats(body)
	case "task_finish":
		j.Lock()
		defer j.Unlock()
		j.NumOfRuns--
		if (body)["reason"] != "success" {
			j.failed = true
		}
		logrus.WithField("remain", j.NumOfRuns).Infoln("Task finished")
		if j.NumOfRuns == 0 {
			go j.Finish()
			return true, nil
		}
	}

	return false, nil
}

func (j *jobProgress) result(runID string) *RunResult {
	j.Lock()
	defer j.Unlock()

	r, ok := j.Runs[runID]
	if !ok {
		r = newRunResult()
		j.Runs[runID] = r
	}
	return r
}

func (j *jobProgress) RecordFailure(body utils.Map) error {
	result, err := json.Marshal(body.Get("result"))
	if err != nil {
		return errors.Wrap(err, "failed to marshal result")
	}

	j.result(body.String("_pansible_task")).
		getSummary(body.String("target")).
		AddFailure(&models.Failure{
			Result: string(result),
		})

	return nil
}

func (j *jobProgress) RecordStats(body utils.Map) {
	target := body.String("target")
	runID := body.String("_pansible_task")

	j.result(runID).getSummary(target).Statistics = &models.Statistic{
		RunID:       runID,
		Target:      target,
		Unreachable: body.Float64("unreachable") == float64(1),
		Failed:      int(body.Float64("failed")),
	}
}

func (j *jobProgress) Stream(ws *websocket.Conn) error {
	t, err := tail.TailFile(j.file.Name(), tail.Config{Follow: true})
	if err != nil {
		return errors.Wrap(err, "failed to tail temp output file")
	}

	defer func() {
		t.Stop()
		t.Cleanup()
	}()

	for {
		select {
		case message := <-t.Lines:
			if err := SendMessage(ws, message.Text); err != nil {
				return errors.Wrap(err, "failed to send message")
			}
		case <-j.done:
			return nil
		}
	}
}

func SendMessage(ws *websocket.Conn, msg interface{}) error {
	s, ok := msg.(string)
	if !ok {
		c, err := json.Marshal(msg)
		if err != nil {
			return errors.Wrap(err, "failed to marshal message")
		}
		s = string(c)
	}

	return errors.Wrap(websocket.Message.Send(ws, s), "failed to send message")
}

func (j *jobProgress) Finish() error {
	close(j.done)
	j.file.Close()
	if err := os.Remove(fmt.Sprintf("output_%s.tmp", j.ID)); err != nil {
		return errors.Wrap(err, "failed to remove temp file")
	}
	return nil
}
