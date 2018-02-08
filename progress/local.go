package progress

import (
	"sync"
	"time"

	"git.elenet.me/yuelong.huang/pansible/models"
	"git.elenet.me/yuelong.huang/pansible/utils"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

type ProgressLocal struct {
	sync.RWMutex
	jobs map[string]*jobProgress
}

func NewLocal() *ProgressLocal {
	return &ProgressLocal{
		jobs: make(map[string]*jobProgress),
	}
}

func (p *ProgressLocal) AddNewJob(jobID string, numOfRuns int) error {
	p.Lock()
	defer p.Unlock()

	if _, exists := p.jobs[jobID]; exists {
		return errors.New("job already exists")
	}

	jp, err := newJobProgress(jobID, time.Now(), numOfRuns)
	if err != nil {
		return errors.Wrap(err, "failed to create job progress")
	}

	p.jobs[jobID] = jp
	return nil
}

func (p *ProgressLocal) get(jobID string) *jobProgress {
	p.RLock()
	jp, _ := p.jobs[jobID]
	p.RUnlock()
	return jp
}

func (p *ProgressLocal) Add(m utils.Map) (done bool, err error) {
	jp := p.get(m.String("_pansible_job"))
	if jp == nil {
		return false, errors.New("job not found")
	}

	return jp.Add(m)
}

func (p *ProgressLocal) Stream(jobID string, ws *websocket.Conn) error {
	jp := p.get(jobID)
	if jp == nil {
		return errors.New("no such job")
	}
	return errors.Wrap(jp.Stream(ws), "stream")
}

func (p *ProgressLocal) Summaries(jobID string) (map[*models.Statistic][]*models.Failure, bool, error) {
	jp := p.get(jobID)
	if jp == nil {
		return nil, true, errors.New("no such job")
	}
	runs := jp.Runs
	summaries := make(map[*models.Statistic][]*models.Failure)
	for _, run := range runs {
		for _, s := range run.Summaries {
			s.Statistics.JobStartTime = jp.StartTime
			summaries[s.Statistics] = s.Failures
		}
	}
	return summaries, jp.failed, nil
}

func (p *ProgressLocal) Clean(jobID string) {
	jp := p.get(jobID)
	if jp == nil {
		return
	}
	select {
	case <-jp.done:
	default:
		jp.Finish()
	}
	p.Lock()
	delete(p.jobs, jobID)
	p.Unlock()
}
