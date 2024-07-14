package coolifygo

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Settings struct {
	ID                         int         `json:"id"`
	CleanupAfterPercentage     int         `json:"cleanup_after_percentage"`
	ConcurrentBuilds           int         `json:"concurrent_builds"`
	DynamicTimeout             int         `json:"dynamic_timeout"`
	ForceDisabled              bool        `json:"force_disabled"`
	IsBuildServer              bool        `json:"is_build_server"`
	IsCloudflareTunnel         bool        `json:"is_cloudflare_tunnel"`
	IsJumpServer               bool        `json:"is_jump_server"`
	IsLogdrainAxiomEnabled     bool        `json:"is_logdrain_axiom_enabled"`
	IsLogdrainCustomEnabled    bool        `json:"is_logdrain_custom_enabled"`
	IsLogdrainHighlightEnabled bool        `json:"is_logdrain_highlight_enabled"`
	IsLogdrainNewrelicEnabled  bool        `json:"is_logdrain_newrelic_enabled"`
	IsMetricsEnabled           bool        `json:"is_metrics_enabled"`
	IsReachable                bool        `json:"is_reachable"`
	IsServerAPIEnabled         bool        `json:"is_server_api_enabled"`
	IsSwarmManager             bool        `json:"is_swarm_manager"`
	IsSwarmWorker              bool        `json:"is_swarm_worker"`
	IsUsable                   bool        `json:"is_usable"`
	LogdrainAxiomAPIKey        interface{} `json:"logdrain_axiom_api_key"`
	LogdrainAxiomDatasetName   interface{} `json:"logdrain_axiom_dataset_name"`
	LogdrainCustomConfig       interface{} `json:"logdrain_custom_config"`
	LogdrainCustomConfigParser interface{} `json:"logdrain_custom_config_parser"`
	LogdrainHighlightProjectID interface{} `json:"logdrain_highlight_project_id"`
	LogdrainNewrelicBaseURI    interface{} `json:"logdrain_newrelic_base_uri"`
	LogdrainNewrelicLicenseKey interface{} `json:"logdrain_newrelic_license_key"`
	MetricsHistoryDays         int         `json:"metrics_history_days"`
	MetricsRefreshRateSeconds  int         `json:"metrics_refresh_rate_seconds"`
	MetricsToken               string      `json:"metrics_token"`
	ServerID                   int         `json:"server_id"`
	WildcardDomain             string      `json:"wildcard_domain"`
	CreatedAt                  time.Time   `json:"created_at"`
	UpdatedAt                  time.Time   `json:"updated_at"`
}

type Server struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	IP          string   `json:"ip"`
	IsReachable bool     `json:"is_reachable"`
	IsUsable    bool     `json:"is_usable"`
	Port        int      `json:"port"`
	Settings    Settings `json:"settings"`
	User        string   `json:"user"`
}

func (c *Client) GetServers() ([]Server, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/servers", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res []Server
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	slog.Info("servers", "res", res)

	return res, nil
}

type Resource struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

func (c *Client) GetServerResources(serverID string) ([]Resource, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/servers/%s/resources", c.baseURL, serverID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res []Resource
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	slog.Info("servers", "res", res)

	return res, nil
}
