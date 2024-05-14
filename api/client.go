package api

type Client interface {
	// Databases 获取数据库列表
	Databases() ([]string, error)

	// Tables 获取表列表
	Tables(database string) ([]string, error)

	// Columns 获取列列表
	Columns(database, table string) ([]ColumnSchema, error)

	// ExecuteQuery 查询
	ExecuteQuery(databaseName, query string) ([]map[string]interface{}, error)

	// Execute 执行
	Execute(database, query string) error

	// Ping 测试连接
	Ping() error

	// TableSchema 获取表结构
	TableSchema(database, table string) (TableSchema, error)
}
