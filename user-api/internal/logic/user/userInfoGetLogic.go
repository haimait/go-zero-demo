package user

import (
	"context"
	"fmt"
	"go-zero-demo/model"
	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoGetLogic {
	return &UserInfoGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoGetLogic) UserInfoGet(req *types.UserInfoGetRequest) (resp *types.UserInfoGetResponse, err error) {
	// todo: add your logic here and delete this line
	//return nil, errorx.NewDefaultError("用户名不存在")
	fmt.Println("req:get-->", req)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Printf("user:%+v", user)
	fmt.Println("err:-->", err)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errors.New("查询用户出错"), "查询用户出错 id:%d , err:%v", req.UserId, err)
	}
	if user == nil {
		return nil, errors.Wrapf(model.ErrNotFound, "id-->:%d", req.UserId)
	}
	var respUser model.User
	_ = copier.Copy(&respUser, user)

	return &types.UserInfoGetResponse{
		UserId:   respUser.Id,
		Nickname: respUser.Name,
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
	//return &types.UserInfoGetResponse{
	//	UserId:   req.UserId,
	//	Nickname: nickname,
	//}, nil
}
