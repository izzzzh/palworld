package container_services

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"palworld/common/docker"
	"strings"
	"time"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContainerLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type EventGenerator struct {
	Channel   chan *LogEvent
	Reader    *bufio.Reader
	LastError error
}

type LogEvent struct {
	Message   any    `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Id        uint32 `json:"id,omitempty"`
}

func NewContainerLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainerLogLogic {
	return &ContainerLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContainerLogLogic) ContainerLog(req *types.ContainerLogReq) (io.ReadCloser, error) {
	// todo: add your logic here and delete this line
	id := req.ID

	cli := docker.ClientPool
	reader, err := cli.ContainerLogs(l.ctx, id, "1000")
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func NewEventGenerator(reader *bufio.Reader) *EventGenerator {
	g := &EventGenerator{
		Channel: make(chan *LogEvent, 100),
		Reader:  reader,
	}
	go g.consume()
	return g
}

func (g *EventGenerator) Next() *LogEvent {
	select {
	case event := <-g.Channel:
		return event
	case <-time.After(50 * time.Millisecond):
		return nil
	}
}

func (g *EventGenerator) consume() {
	for {
		message, readerError := g.Reader.ReadString('\n')
		if readerError != nil {
			g.LastError = readerError
			break
		}
		if message != "" {
			local, _ := time.LoadLocation("Asia/Shanghai")
			logEvent := &LogEvent{
				Message:   message,
				Timestamp: time.Now().In(local).Unix(),
			}
			if index := strings.IndexAny(message, " "); index != -1 {
				logId := message[:index]
				if timestamp, err := time.Parse(time.RFC3339Nano, logId); err == nil {
					logEvent.Timestamp = timestamp.UnixMilli()
				}
				message = strings.TrimSuffix(message[index+1:], "\n")
				logEvent.Message = message
				if strings.HasPrefix(message, "{") && strings.HasSuffix(message, "}") {
					var data map[string]interface{}
					if err := json.Unmarshal([]byte(message), &data); err != nil {
						logx.Info("json unmarshal error while streaming %v", err.Error())
					} else {
						logEvent.Message = data
					}
				}
			}
			g.Channel <- logEvent
		}
	}
}
