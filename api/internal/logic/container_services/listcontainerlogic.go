package container_services

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
	"palworld/thrid/docker"
	"strings"
)

type ListContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContainerLogic {
	return &ListContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContainerLogic) ListContainer() (*types.ListContainerResp, error) {

	var ret = &types.ListContainerResp{}

	cli := docker.ClientPool

	listCon, err := cli.ListContainers()
	if err != nil {
		panic(err)
	}

	containers := make([]types.Container, 0)
	for _, val := range listCon {
		if strings.HasPrefix(val.Image, "sha256") {
			continue
		}
		containers = append(containers, types.Container{
			ID:      val.ID,
			Name:    val.Name,
			Image:   val.Image,
			State:   val.State,
			Status:  val.Status,
			Created: val.Created,
			Health:  val.Health,
		})
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = containers
	return ret, nil
}
