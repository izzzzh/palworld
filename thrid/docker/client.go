package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	"io"
	"regexp"
	"strings"
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
	ContainerStats(context.Context, string, chan<- ContainerStat) error
	Ping(context.Context) (types.Ping, error)
}

func NewClient() (Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
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
			ID:      val.ID[:12],
			Names:   val.Names,
			Name:    strings.TrimPrefix(val.Names[0], "/"),
			Image:   val.Image,
			ImageID: val.ImageID,
			Command: val.Command,
			Created: val.Created,
			State:   val.State,
			Status:  val.Status,
			Health:  findBetweenParentheses(val.Status),
		})
	}
	return ret, nil
}

func (d dockerClient) ContainerLogs(ctx context.Context, containerID string, tail string) (io.ReadCloser, error) {
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Tail:       "300",
		Timestamps: true,
	}
	reader, err := d.client.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (d dockerClient) ContainerStats(ctx context.Context, containerID string, ch chan<- ContainerStat) error {
	return nil
}
func (d dockerClient) Ping(ctx context.Context) (types.Ping, error) {
	return types.Ping{}, nil
}

var PARENTHESIS_RE = regexp.MustCompile(`\(([a-zA-Z]+)\)`)

func findBetweenParentheses(s string) string {
	if results := PARENTHESIS_RE.FindStringSubmatch(s); results != nil {
		return results[1]
	}
	return ""
}
