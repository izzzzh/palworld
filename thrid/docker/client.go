package docker

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
)

type dockerProxy interface {
	ContainerList(context.Context, types.ContainerListOptions) ([]types.Container, error)
	ContainerLogs(context.Context, string, types.ContainerLogsOptions) (io.ReadCloser, error)
	Events(context.Context, types.EventsOptions) (<-chan events.Message, <-chan error)
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
	ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
	Ping(ctx context.Context) (types.Ping, error)
}

type dockerClient struct {
	client dockerProxy
}

type Client interface {
	ListContainers() ([]Container, error)
	ContainerLogs(context.Context, string, string) (io.ReadCloser, error)
	ContainerLogsBetweenDates(context.Context, string, time.Time, time.Time) (io.ReadCloser, error)
	ContainerStats(context.Context, string, chan<- ContainerStat) error
	Ping(context.Context) (types.Ping, error)
}

func NewClient() (Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &dockerClient{client: cli}, nil
}

func (d dockerClient) ListContainers() ([]Container, error) {
	var ret []Container
	containers, err := d.client.ContainerList(context.Background(),
		types.ContainerListOptions{All: true},
	)
	if err != nil {
		return nil, err
	}

	for _, val := range containers {
		ret = append(ret, Container{
			ID:      val.ID,
			Image:   val.Image,
			Command: val.Command,
			Created: val.Created,
			State:   val.State,
			Status:  val.Status,
			Names:   val.Names,
		})
	}
	return ret, nil
}

func (d dockerClient) ContainerLogs(ctx context.Context, containerID string, tail string) (io.ReadCloser, error) {
	return d.client.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       tail,
	})
}

func (d dockerClient) ContainerLogsBetweenDates(ctx context.Context, containerID string, from time.Time, to time.Time) (io.ReadCloser, error) {
	return nil, nil
}

func (d dockerClient) ContainerStats(ctx context.Context, containerID string, ch chan<- ContainerStat) error {
	return nil
}
func (d dockerClient) Ping(ctx context.Context) (types.Ping, error) {
	return types.Ping{}, nil
}
