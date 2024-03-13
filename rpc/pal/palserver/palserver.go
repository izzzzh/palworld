// Code generated by goctl. DO NOT EDIT.
// Source: pal.proto

package palserver

import (
	"context"

	"palworld/rpc/pal/pb/pal"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Ability     = pal.Ability
	GetPalReq   = pal.GetPalReq
	GetPalResp  = pal.GetPalResp
	ListPal     = pal.ListPal
	ListPalReq  = pal.ListPalReq
	ListPalResp = pal.ListPalResp
	Pal         = pal.Pal
	Skill       = pal.Skill

	PalServer interface {
		GetPal(ctx context.Context, in *GetPalReq, opts ...grpc.CallOption) (*GetPalResp, error)
		ListPal(ctx context.Context, in *ListPalReq, opts ...grpc.CallOption) (*ListPalResp, error)
	}

	defaultPalServer struct {
		cli zrpc.Client
	}
)

func NewPalServer(cli zrpc.Client) PalServer {
	return &defaultPalServer{
		cli: cli,
	}
}

func (m *defaultPalServer) GetPal(ctx context.Context, in *GetPalReq, opts ...grpc.CallOption) (*GetPalResp, error) {
	client := pal.NewPalServerClient(m.cli.Conn())
	return client.GetPal(ctx, in, opts...)
}

func (m *defaultPalServer) ListPal(ctx context.Context, in *ListPalReq, opts ...grpc.CallOption) (*ListPalResp, error) {
	client := pal.NewPalServerClient(m.cli.Conn())
	return client.ListPal(ctx, in, opts...)
}
