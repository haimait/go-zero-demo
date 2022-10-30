package no_cache_model

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

//统一model 执行接口
type Executable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var UserAuthTypeSystem string = "system"  //平台内部
var UserAuthTypeSmallWX string = "wxMini" //微信小程序
