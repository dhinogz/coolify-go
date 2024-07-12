package coolifygo

import (
	"fmt"
	"net/http"
)

type Deployment struct {
	ID               int    `json:"id"`
	ApplicationID    string `json:"application_id"`
	DeploymentUUID   string `json:"deployment_uuid"`
	PullRequestID    int    `json:"pull_request_id"`
	ForceRebuild     bool   `json:"force_rebuild"`
	Commit           string `json:"commit"`
	Status           string `json:"status"`
	IsWebhook        bool   `json:"is_webhook"`
	IsAPI            bool   `json:"is_api"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	Logs             string `json:"logs"`
	CurrentProcessID string `json:"current_process_id"`
	RestartOnly      bool   `json:"restart_only"`
	GitType          string `json:"git_type"`
	ServerID         int    `json:"server_id"`
	ApplicationName  string `json:"application_name"`
	ServerName       string `json:"server_name"`
	DeploymentURL    string `json:"deployment_url"`
	DestinationID    string `json:"destination_id"`
	OnlyThisServer   bool   `json:"only_this_server"`
	Rollback         bool   `json:"rollback"`
	CommitMessage    string `json:"commit_message"`
}

func (c *Client) GetDeployments() ([]Deployment, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res []Deployment
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
