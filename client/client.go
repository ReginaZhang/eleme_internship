package client

import (
	"fmt"
	"net/http"
	"time"

	"git.elenet.me/yuelong.huang/pansible/models"
	"github.com/franela/goreq"
	"github.com/pkg/errors"
)

type Client struct {
	Master string
	Token  string
}

func New(master string, token string) *Client {
	return &Client{
		Master: master,
		Token:  token,
	}
}

func (c *Client) req(method, path string) goreq.Request {
	return goreq.Request{
		Timeout: time.Second,
		Method:  method,
		Uri:     fmt.Sprintf("http://%s/runner/%s", c.Master, path),
	}.WithHeader("Authorization", fmt.Sprintf("Bearer %s", c.Token))
}

func (c *Client) get(path string) (*goreq.Response, error) {
	req := c.req(http.MethodGet, path)
	return req.Do()
}

func httpError(resp *goreq.Response, err error) error {
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("request failed with rc %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) GetInventoryByRunID(runID string) (*models.Inventory, error) {

	resp, err := c.get(fmt.Sprintf("inventory?run=%d", runID))
	if err := httpError(resp, err); err != nil {
		return nil, errors.Wrap(err, "failed to get inventory")
	}

	defer resp.Body.Close()

	inv := models.Inventory{}
	if err := resp.Body.FromJsonTo(&inv); err != nil {
		return nil, errors.Wrap(err, "failed to parse body")
	}

	return &inv, nil
}

func (c *Client) GetPlaybookByJobID(jobID string) (*models.Playbook, error) {

	resp, err := c.get(fmt.Sprintf("playbook/%s", jobID))
	if err := httpError(resp, err); err != nil {
		return nil, errors.Wrap(err, "failed to get playbook")
	}

	defer resp.Body.Close()

	playbook := models.Playbook{}
	if err := resp.Body.FromJsonTo(&playbook); err != nil {
		return nil, errors.Wrap(err, "failed to parse body")
	}

	return &playbook, nil
}

func (c *Client) GetRun(uuid string) (*models.Run, error) {

	resp, err := c.get(fmt.Sprintf("run/%s", uuid))
	if err := httpError(resp, err); err != nil {
		return nil, errors.Wrap(err, "failed to get run")
	}

	defer resp.Body.Close()

	run := models.Run{}
	if err := resp.Body.FromJsonTo(&run); err != nil {
		return nil, errors.Wrap(err, "failed to parse body")
	}

	return &run, nil
}
