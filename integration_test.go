package coolifygo

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	COOLIFY_API_KEY = "api-key"
	COOLIFY_ADDR    = "coolify-addr"
)

func TestGetProjects(t *testing.T) {
	c := NewClient(COOLIFY_ADDR, COOLIFY_API_KEY)

	res, err := c.GetProjects()
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		slog.Info("projects", "res", res)
	}
}
