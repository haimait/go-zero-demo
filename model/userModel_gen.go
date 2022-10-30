// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"go-zero-demo/common/globalkey"
	"go-zero-demo/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
)

type (
	userModel interface {
		//根据主键查询一条数据，走缓存
		FindOne(ctx context.Context, id int64) (*User, error)
		//根据唯一索引查询一条数据，走缓存
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)

		//删除数据
		Delete(ctx context.Context, session sqlx.Session, id int64) error

		//新增数据
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)

		//软删除数据
		DeleteSoft(ctx context.Context, session sqlx.Session, data *User) error

		//更新数据
		Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)

		//更新数据，使用乐观锁
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *User) error

		//根据条件查询一条数据，不走缓存
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*User, error)

		//sum某个字段
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)

		//根据条件统计条数
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)

		//查询所有数据不分页
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error)

		//根据页码分页查询分页数据
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, error)

		//根据id倒序分页查询分页数据
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*User, error)

		//根据id升序分页查询分页数据
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*User, error)

		//暴露给logic，开启事务
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error

		//暴露给logic，查询数据的builder
		RowBuilder() squirrel.SelectBuilder

		//暴露给logic，查询count的builder
		CountBuilder(field string) squirrel.SelectBuilder

		//暴露给logic，查询sum的builder
		SumBuilder(field string) squirrel.SelectBuilder
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id        int64     `db:"id"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		DeletedAt time.Time `db:"deleted_at"`
		Version   int64     `db:"version"` // 版本号
		Mobile    string    `db:"mobile"`
		Password  string    `db:"password"`
		Nickname  string    `db:"nickname"`
		Sex       int64     `db:"sex"` // 性别 0:男 1:女
		Avatar    string    `db:"avatar"`
		Info      string    `db:"info"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {

	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

//软删除数据
func (m *defaultUserModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *User) error {
	data.DeletedAt = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"), "UserModel delete err : %+v", err)
	}
	return nil
}

//根据主键查询一条数据，走缓存
func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//根据唯一索引查询一条数据，走缓存
func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	data.DeletedAt = time.Unix(0, 0)

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.DeletedAt, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
	}
	return m.conn.ExecCtx(ctx, query, data.DeletedAt, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)

}

//暴露给logic开启事务
func (m *defaultUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, newData *User) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	if session != nil {
		return session.ExecCtx(ctx, query, newData.DeletedAt, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.Id)
	}
	return m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.Id)
}

//乐观锁修改数据 ,推荐使用
func (m *defaultUserModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *User) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, userRowsWithPlaceHolder)
	if session != nil {
		sqlResult, err = session.ExecCtx(ctx, query, newData.DeletedAt, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.Id, oldVersion)
	} else {
		sqlResult, err = m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.Id, oldVersion)
	}

	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil

}

//根据条件查询一条数据
func (m *defaultUserModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*User, error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp User

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

//统计某个字段总和
func (m *defaultUserModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

//根据某个字段查询数据数量
func (m *defaultUserModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64

	err = m.conn.QueryRowCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

//查询所有数据
func (m *defaultUserModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照页码分页查询数据
func (m *defaultUserModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id倒序分页查询数据，不支持排序
func (m *defaultUserModel) FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*User, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *defaultUserModel) FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*User, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//暴露给logic查询数据构建条件使用的builder
func (m *defaultUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userRows).From(m.table)
}

//暴露给logic查询count构建条件使用的builder
func (m *defaultUserModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

//暴露给logic查询构建条件使用的builder
func (m *defaultUserModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
