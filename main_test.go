package main

import (
	"context"
	"github.com/go-integration-tests/testContainers/container"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	containerName = "my-container"
	dockerfile    = "Dockerfile"
	composePath   = "./docker-compose.yml"
)

func TestContainerService(t *testing.T) {
	currDir, err := os.Getwd()
	require.NoError(t, err)

	ctx := context.Background()
	c, err := container.NewContainer(ctx, containerName, currDir, dockerfile)
	time.Sleep(10 * time.Second)
	require.NoError(t, err)
	require.NotNil(t, c)
	require.NotNil(t, c.Container)

	// clean up
	require.NoError(t, c.Terminate(ctx))
}

func TestComposeServices(t *testing.T) {
	identifier := strings.ToLower(uuid.New().String())
	compose := testcontainers.NewLocalDockerCompose([]string{composePath}, identifier)
	assert.NotNil(t, compose)
	time.Sleep(10 * time.Second)
	destroyFn := func() {
		err := compose.Down()
		require.NoError(t, err.Error)
	}
	defer destroyFn()
	err := compose.WithCommand([]string{"up", "-d"}).Invoke()
	require.NoError(t, err.Error)
}
