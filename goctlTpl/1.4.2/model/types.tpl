
type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}

		//新增数据
		Insert(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) (sql.Result,error)

		//软删除数据
		DeleteSoft(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) error

		//更新数据
		Update(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}})  (sql.Result, error)

		//更新数据，使用乐观锁
		UpdateWithVersion(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) error

		//根据条件查询一条数据，不走缓存
		FindOneByQuery(ctx context.Context,rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}},error)

		//sum某个字段
		FindSum(ctx context.Context,sumBuilder squirrel.SelectBuilder) (float64,error)

		//根据条件统计条数
		FindCount(ctx context.Context,countBuilder squirrel.SelectBuilder) (int64,error)

		//查询所有数据不分页
		FindAll(ctx context.Context,rowBuilder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error)

		//根据页码分页查询分页数据
		FindPageListByPage(ctx context.Context,rowBuilder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error)

		//根据id倒序分页查询分页数据
		FindPageListByIdDESC(ctx context.Context,rowBuilder squirrel.SelectBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)

		//根据id升序分页查询分页数据
		FindPageListByIdASC(ctx context.Context,rowBuilder squirrel.SelectBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)

		//暴露给logic，开启事务
		Trans(ctx context.Context,fn func(context context.Context,session sqlx.Session) error) error

		//暴露给logic，查询数据的builder
		RowBuilder() squirrel.SelectBuilder

		//暴露给logic，查询count的builder
		CountBuilder(field string) squirrel.SelectBuilder

		//暴露给logic，查询sum的builder
		SumBuilder(field string) squirrel.SelectBuilder
}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
