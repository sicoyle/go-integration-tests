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
	sleepTime     = 10 * time.Second
)

func TestContainerService(t *testing.T) {
	// setup
	currDir, err := os.Getwd()
	require.NoError(t, err)
	ctx := context.Background()

	// start container for test
	c, err := container.NewContainer(ctx, containerName, currDir, dockerfile)
	time.Sleep(sleepTime)
	require.NoError(t, err)
	require.NotNil(t, c)

	// clean up
	require.NoError(t, c.Terminate(ctx))
}

func TestComposeServices(t *testing.T) {
	// setup
	identifier := strings.ToLower(uuid.New().String())

	// create compose
	compose := testcontainers.NewLocalDockerCompose([]string{composePath}, identifier)
	assert.NotNil(t, compose)
	time.Sleep(sleepTime)

	// defer cleanup
	destroyFn := func() {
		err := compose.Down()
		require.NoError(t, err.Error)
	}
	defer destroyFn()

	// start services for test
	err := compose.WithCommand([]string{"up", "-d"}).Invoke()
	require.NoError(t, err.Error)
}
