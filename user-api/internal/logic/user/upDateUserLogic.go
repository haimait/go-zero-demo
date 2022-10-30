package user

import (
	"context"

	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateUserLogic {
	return &UpDateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateUserLogic) UpDateUser(req *types.UserUpdateRequest) (resp *types.UserUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	// 这里写111
	return
}
