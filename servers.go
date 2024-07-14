package coolifygo

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	ID                              int       `json:"id"`
	RepositoryProjectID             int       `json:"repository_project_id"`
	UUID                            string    `json:"uuid"`
	Name                            string    `json:"name"`
	Fqdn                            string    `json:"fqdn"`
	ConfigHash                      string    `json:"config_hash"`
	GitRepository                   string    `json:"git_repository"`
	GitBranch                       string    `json:"git_branch"`
	GitCommitSha                    string    `json:"git_commit_sha"`
	GitFullURL                      string    `json:"git_full_url"`
	DockerRegistryImageName         string    `json:"docker_registry_image_name"`
	DockerRegistryImageTag          string    `json:"docker_registry_image_tag"`
	BuildPack                       string    `json:"build_pack"`
	StaticImage                     string    `json:"static_image"`
	InstallCommand                  string    `json:"install_command"`
	BuildCommand                    string    `json:"build_command"`
	StartCommand                    string    `json:"start_command"`
	PortsExposes                    string    `json:"ports_exposes"`
	PortsMappings                   string    `json:"ports_mappings"`
	BaseDirectory                   string    `json:"base_directory"`
	PublishDirectory                string    `json:"publish_directory"`
	HealthCheckPath                 string    `json:"health_check_path"`
	HealthCheckPort                 string    `json:"health_check_port"`
	HealthCheckHost                 string    `json:"health_check_host"`
	HealthCheckMethod               string    `json:"health_check_method"`
	HealthCheckReturnCode           int       `json:"health_check_return_code"`
	HealthCheckScheme               string    `json:"health_check_scheme"`
	HealthCheckResponseText         string    `json:"health_check_response_text"`
	HealthCheckInterval             int       `json:"health_check_interval"`
	HealthCheckTimeout              int       `json:"health_check_timeout"`
	HealthCheckRetries              int       `json:"health_check_retries"`
	HealthCheckStartPeriod          int       `json:"health_check_start_period"`
	LimitsMemory                    string    `json:"limits_memory"`
	LimitsMemorySwap                string    `json:"limits_memory_swap"`
	LimitsMemorySwappiness          int       `json:"limits_memory_swappiness"`
	LimitsMemoryReservation         string    `json:"limits_memory_reservation"`
	LimitsCpus                      string    `json:"limits_cpus"`
	LimitsCpuset                    string    `json:"limits_cpuset"`
	LimitsCPUShares                 int       `json:"limits_cpu_shares"`
	Status                          string    `json:"status"`
	PreviewURLTemplate              string    `json:"preview_url_template"`
	DestinationType                 string    `json:"destination_type"`
	DestinationID                   int       `json:"destination_id"`
	SourceType                      string    `json:"source_type"`
	SourceID                        int       `json:"source_id"`
	PrivateKeyID                    int       `json:"private_key_id"`
	EnvironmentID                   int       `json:"environment_id"`
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
	Description                     string    `json:"description"`
	Dockerfile                      string    `json:"dockerfile"`
	HealthCheckEnabled              bool      `json:"health_check_enabled"`
	DockerfileLocation              string    `json:"dockerfile_location"`
	CustomLabels                    string    `json:"custom_labels"`
	DockerfileTargetBuild           string    `json:"dockerfile_target_build"`
	ManualWebhookSecretGithub       string    `json:"manual_webhook_secret_github"`
	ManualWebhookSecretGitlab       string    `json:"manual_webhook_secret_gitlab"`
	DockerComposeLocation           string    `json:"docker_compose_location"`
	DockerCompose                   string    `json:"docker_compose"`
	DockerComposeRaw                string    `json:"docker_compose_raw"`
	DockerComposeDomains            string    `json:"docker_compose_domains"`
	DeletedAt                       time.Time `json:"deleted_at"`
	DockerComposeCustomStartCommand string    `json:"docker_compose_custom_start_command"`
	DockerComposeCustomBuildCommand string    `json:"docker_compose_custom_build_command"`
	SwarmReplicas                   int       `json:"swarm_replicas"`
	SwarmPlacementConstraints       string    `json:"swarm_placement_constraints"`
	ManualWebhookSecretBitbucket    string    `json:"manual_webhook_secret_bitbucket"`
	CustomDockerRunOptions          string    `json:"custom_docker_run_options"`
	PostDeploymentCommand           string    `json:"post_deployment_command"`
	PostDeploymentCommandContainer  string    `json:"post_deployment_command_container"`
	PreDeploymentCommand            string    `json:"pre_deployment_command"`
	PreDeploymentCommandContainer   string    `json:"pre_deployment_command_container"`
	WatchPaths                      string    `json:"watch_paths"`
	CustomHealthcheckFound          bool      `json:"custom_healthcheck_found"`
	ManualWebhookSecretGitea        string    `json:"manual_webhook_secret_gitea"`
	Redirect                        string    `json:"redirect"`
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
