package user

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-demo/user-rpc/usercenter"

	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoRpcGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoRpcGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoRpcGetLogic {
	return &UserInfoRpcGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoRpcGetLogic) UserInfoRpcGet(req *types.UserInfoGetRpcRequest) (resp *types.UserInfoGetRpcResponse, err error) {
	userInfoResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id:       req.UserId,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	//return nil, errorx.NewCodeError(4001, "用户密码不正确")
	//return nil, errorx.NewDefaultError("用户密码不正确")
	//return nil, errors.New("用户密码不正确")
	var userInfo types.UserInfoGetRpcResponse
	_ = copier.Copy(&userInfo, userInfoResp.User)

	return &userInfo, nil
}
