package container_services

import (
	"context"
	"net/http"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
	"palworld/thrid/docker"

	"github.com/zeromicro/go-zero/core/logx"
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
	var ret *types.ListContainerResp

	cli, err := docker.NewClient()
	listCon, err := cli.ListContainers()
	if err != nil {
		ret.Code = http.StatusInternalServerError
		ret.Message = err.Error()
		return ret, nil
	}

	containers := make([]types.Container, 0)
	for _, val := range listCon {
		containers = append(containers, types.Container{
			ID:    val.ID,
			Name:  val.Names[0],
			Image: val.Image,
		})
	}

	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = containers
	return ret, nil
}
