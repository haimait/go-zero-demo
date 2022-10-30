
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}



//乐观锁修改数据 ,推荐使用
func (m *default{{.upperStartCamelObject}}Model) UpdateWithVersion(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	{{if .withCache}}{{.keys}}
	sqlResult,err =  m.ExecCtx(ctx,func(ctx context.Context,conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}
	return conn.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		sqlResult,err  =  session.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}else{
		sqlResult,err  =  m.conn.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}
	{{end}}
	if err != nil {
		return err
	}
	updateCount , err := sqlResult.RowsAffected()
	if err != nil{
		return err
	}
	if updateCount == 0 {
	return  xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil

}

//根据条件查询一条数据
func (m *default{{.upperStartCamelObject}}Model) FindOneByQuery(ctx context.Context,rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}},error) {

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp {{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
	return &resp, nil
	default:
	return nil, err
	}
}


//统计某个字段总和
func (m *default{{.upperStartCamelObject}}Model) FindSum(ctx context.Context,sumBuilder squirrel.SelectBuilder) (float64,error) {

	query, values, err := sumBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	{{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
		err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}


//根据某个字段查询数据数量
func (m *default{{.upperStartCamelObject}}Model) FindCount(ctx context.Context,countBuilder squirrel.SelectBuilder) (int64,error) {

	query, values, err := countBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
	return 0, err
	}

	var resp int64
	{{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}


//查询所有数据
func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,rowBuilder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error) {

	if orderBy == ""{
		rowBuilder = rowBuilder.OrderBy("id DESC")
	}else{
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}


//按照页码分页查询数据
func (m *default{{.upperStartCamelObject}}Model) FindPageListByPage(ctx context.Context,rowBuilder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error) {

	if orderBy == ""{
		rowBuilder = rowBuilder.OrderBy("id DESC")
	}else{
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1{
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}


//按照id倒序分页查询数据，不支持排序
func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdDESC(ctx context.Context,rowBuilder squirrel.SelectBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? " , preMinId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
		err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdASC(ctx context.Context,rowBuilder squirrel.SelectBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)  {


	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? " , preMaxId)
	}

	query, values, err := rowBuilder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
	err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}



//暴露给logic查询数据构建条件使用的builder
func (m *default{{.upperStartCamelObject}}Model) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select({{.lowerStartCamelObject}}Rows).From(m.table)
}

//暴露给logic查询count构建条件使用的builder
func (m *default{{.upperStartCamelObject}}Model) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT("+field+")").From(m.table)
}

//暴露给logic查询构建条件使用的builder
func (m *default{{.upperStartCamelObject}}Model) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM("+field+"),0)").From(m.table)
}
