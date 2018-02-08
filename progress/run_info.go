package progress

import (
	"sync"
	"time"

	"git.elenet.me/yuelong.huang/pansible/models"
)

type RunResult struct {
	sync.RWMutex
	Summaries map[string]*Summary
	StartTime time.Time
}

type Summary struct {
	Failures   []*models.Failure
	Statistics *models.Statistic
}

func (s *Summary) AddFailure(f *models.Failure) {
	s.Failures = append(s.Failures, f)
}

func newRunResult() *RunResult {
	return &RunResult{
		Summaries: make(map[string]*Summary),
	}
}

func (r *RunResult) getSummary(target string) *Summary {
	r.Lock()
	defer r.Unlock()

	s, ok := r.Summaries[target]
	if !ok {
		s = &Summary{
			Failures:   []*models.Failure{},
			Statistics: &models.Statistic{},
		}
		r.Summaries[target] = s
	}

	return s
}
