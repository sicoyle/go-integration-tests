package container

import (
	"bytes"
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"io/ioutil"
	"time"
)

const (
	helloMsg = "Hello gophers!\n"
	port     = "8081"
)

type myContainer struct {
	testcontainers.Container
	URI string
}

func NewContainer(ctx context.Context, containerName, dockerContextDir, dockerFileName string) (*myContainer, error) {
	req := testcontainers.ContainerRequest{
		Name: containerName,
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    dockerContextDir,
			Dockerfile: dockerFileName,
		},
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", port)},
		WaitingFor: wait.ForHTTP("/").
			WithStartupTimeout(10 * time.Second).WithPort(port).
			WithResponseMatcher(func(body io.Reader) bool {
				data, err := ioutil.ReadAll(body)
				if err != nil {
					return false
				}
				return bytes.Equal(data, []byte(helloMsg))
			}),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}
	mappedPort, err := container.MappedPort(ctx, port)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("http://%s:%s", ip, mappedPort.Port())
	return &myContainer{Container: container, URI: uri}, nil
}
