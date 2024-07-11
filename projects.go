package coolifygo

import (
	"fmt"
	"log/slog"
	"net/http"
)

// Project represents a project in the API
type Project struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	// Environments []Environment `json:"environments"`
}

// Environment represents an environment in the API
type Environment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProjectID   int    `json:"project_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Description string `json:"description"`
}

// GetFaces returns trained faces and their images
func (c *Client) GetProjects() ([]Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res []Project
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	slog.Info("projects", "res", res)

	return res, nil
}
