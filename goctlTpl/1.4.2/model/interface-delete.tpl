//删除数据
Delete(ctx context.Context,session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error
