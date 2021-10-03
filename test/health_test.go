//go:build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func TestHealthEndpoint(t *testing.T) {
	log.Info("running e2e test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}
