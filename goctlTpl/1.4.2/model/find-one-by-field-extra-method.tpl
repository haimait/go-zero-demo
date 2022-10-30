//格式化缓存key
func (m *default{{.upperStartCamelObject}}Model) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", {{.primaryKeyLeft}}, primary)
}
//根据主键去db查询一条数据
func (m *default{{.upperStartCamelObject}}Model) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where {{.originalPrimaryField}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table )
	return conn.QueryRowCtx(ctx, v, query, primary)
}

//----------------------------------------其他自定义方法，从此处开始写,此处上方不要写自定义方法----------------------------------------