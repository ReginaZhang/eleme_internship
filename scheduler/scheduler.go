package scheduler

import (
	"strings"

	"git.elenet.me/yuelong.huang/pansible/models"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Scheduler interface {
	Schedule(inv *models.Inventory) ([]*models.Run, error)
}

type LocalScheduler struct{}
type ConstantScheduler struct {
	size int
}

func NewLocal() *LocalScheduler {
	return &LocalScheduler{}
}

func NewConstant(size int) *ConstantScheduler {
	return &ConstantScheduler{
		size: size,
	}
}

func (s *LocalScheduler) Schedule(inv *models.Inventory) ([]*models.Run, error) {

	var runs []*models.Run

	hostsMap, err := getHostsMap(inv)
	if err != nil {
		return nil, errors.Wrap(err, "failed to schedule runs")
	}
	if hostsMap == nil {
		return nil, errors.Wrap(err, "hosts field is empty")
	}

	h, err := yaml.Marshal(hostsMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal hosts for run")
	}

	run := models.Run{
		Hosts: string(h),
		Vars:  inv.Vars,
		Env:   inv.Env,
	}

	runs = append(runs, &run)

	return runs, nil
}

func (s *ConstantScheduler) Schedule(inv *models.Inventory) ([]*models.Run, error) {
	var runs []*models.Run

	hostsMap, err := getHostsMap(inv)
	if err != nil {
		return nil, errors.Wrap(err, "failed to schedule runs")
	}
	if hostsMap == nil {
		return nil, errors.Wrap(err, "hosts field is empty")
	}

	h, err := yaml.Marshal(hostsMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal hosts for run")
	}

	hostList := make([]string, 0)
	hostsMapToList(hostsMap, &hostList)

	for i := 0; i < len(hostList); i += s.size {

		var limit string
		if i+s.size >= len(hostList) {
			limit = strings.Join(hostList[i:], "\n")
		} else {
			limit = strings.Join(hostList[i:(i+s.size)], "\n")
		}

		run := models.Run{
			Hosts: string(h),
			Vars:  inv.Vars,
			Env:   inv.Env,
			Limit: limit,
		}

		runs = append(runs, &run)
	}

	return runs, nil
}

func getHostsMap(inv *models.Inventory) (map[interface{}]interface{}, error) {
	hostsMap := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(inv.Hosts), &hostsMap)
	if err != nil {
		return nil, err
	}

	_ = parseHostsData(&hostsMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse hosts map")
	}

	return hostsMap, nil
}

func hostsMapToList(m map[interface{}]interface{}, list *[]string) {
	for key, value := range m {
		if key == "hosts" {
			original, _ := value.([]string)
			*list = append(*list, original...)
		} else {
			original, ok := value.(map[interface{}]interface{})
			if ok {
				hostsMapToList(original, list)
			}
		}
	}
}
