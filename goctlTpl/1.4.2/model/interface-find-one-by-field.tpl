//根据唯一索引查询一条数据，走缓存
FindOneBy{{.upperField}}(ctx context.Context, {{.in}}) (*{{.upperStartCamelObject}}, error)
