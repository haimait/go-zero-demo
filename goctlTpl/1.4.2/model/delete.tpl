func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}
{{end}}	{{.keys}}
	_, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		if session!=nil{
			return session.ExecCtx(ctx,query, {{.lowerStartCamelPrimaryKey}})
		}
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}
        query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		if session!=nil{
			_,err:= session.ExecCtx(ctx,query, {{.lowerStartCamelPrimaryKey}})
			return err
		}
		_,err:=m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
		return err
}


//软删除数据
func (m *default{{.upperStartCamelObject}}Model) DeleteSoft(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) error {
	data.DeletedAt = time.Now()
	if err:= m.UpdateWithVersion(ctx,session, data);err!= nil{
		return errors.Wrapf(xerr.NewErrMsg("删除数据失败"),"{{.upperStartCamelObject}}Model delete err : %+v",err)
	}
	return nil
}


