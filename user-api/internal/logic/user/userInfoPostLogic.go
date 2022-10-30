package user

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-zero-demo/model"

	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoPostLogic {
	return &UserInfoPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoPostLogic) UserInfoPost(req *types.UserInfoPostRequest) (resp *types.UserInfoPostResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("req:post-->", req)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("user:", user)
	fmt.Println("err:-->", err)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errors.New("查询用户出错"), "查询用户出错 id:%d , err:%v", req.UserId, err)
	}
	if user == nil {
		return nil, errors.Wrapf(model.ErrNotFound, "id:%d", req.UserId)
	}
	var respUser model.User
	_ = copier.Copy(&respUser, user)

	return &types.UserInfoPostResponse{
		UserId:   respUser.Id,
		Nickname: respUser.Nickname,
	}, nil

	//users := map[int64]string{
	//	1: "张三",
	//	2: "李四",
	//}
	//nickname := "unknown"
	//fmt.Println("resp:", resp)
	//if name, ok := users[req.UserId]; ok {
	//	nickname = name
	//}
	//fmt.Println("resp:", resp)
	//return &types.UserInfoPostResponse{
	//	UserId:   req.UserId,
	//	Nickname: nickname,
	//}, nil
	//return
}
