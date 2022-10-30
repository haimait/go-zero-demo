package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-zero-demo/model"
	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	var user model.User
	_ = copier.Copy(&user, req)
	insert, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, errors.Wrapf(err, "UserModel.Insert failed err:%v,user:%+v", err, user)
	}
	lastId, err := insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(err, "insertResult.LastInsertId err:%v,user:%+v", err, user)
	}
	var respData types.RegisterResp
	respData.Id = lastId
	//_ = copier.Copy(&resp, registerResp)

	return &respData, nil
}
