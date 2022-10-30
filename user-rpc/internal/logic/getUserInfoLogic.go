package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-zero-demo/model"

	"go-zero-demo/user-rpc/internal/svc"
	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println("in:get-->", in)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	fmt.Printf("user:%+v", user)
	fmt.Println("err:-->", err)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errors.New("查询用户出错"), "查询用户出错 id:%d , err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(model.ErrNotFound, "id-->:%d", in.Id)
	}

	//var respUser model.User
	var respUser pb.User
	_ = copier.Copy(&respUser, user)

	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil

	//return &pb.GetUserInfoResp{}, nil
}
