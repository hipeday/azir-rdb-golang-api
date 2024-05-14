package api

// DataSourceType 数据源类型
type DataSourceType string

const (
	MySQL DataSourceType = "mysql"
	// 其他数据源类型
)

// ConnectionInfo 连接信息
type ConnectionInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// TableSchema defines the schema of a database table
type TableSchema struct {
	Columns []ColumnSchema
}

// ColumnSchema defines the schema of a column in a database table
type ColumnSchema struct {
	Name         string
	DataType     string
	IsNullable   bool
	IsPrimaryKey bool
}
