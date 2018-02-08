package api

import (
	"git.elenet.me/yuelong.huang/pansible/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
)

func (s *Server) FindJobByID(exec boil.Executor, id int64) (*models.Job, error) {
	job, err := models.FindJob(exec, id)
	return job, err
}

func ValidateJob(job *models.Job) error {

	if job.PlaybookID == 0 {
		return errors.New("Missing required parameter 'PlaybookID'")
	}

	if job.InventoryID == 0 {
		return errors.New("Missing required parameter 'InventoryID'")
	}

	if job.Env == "" {
		return errors.New("Missing required parameter 'Env'")
	}

	if err := validateJobEnv(job); err != nil {
		return errors.Wrap(err, "'Env' validation failed")
	}

	return nil
}

func validateJobEnv(job *models.Job) error {
	// TODO
	return nil
}
