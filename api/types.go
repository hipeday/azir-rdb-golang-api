package api

// DataSourceType 数据源类型
type DataSourceType string

const (
	MySQL DataSourceType = "mysql"
	// 其他数据源类型
)

// ConnectionInfo 连接信息
type ConnectionInfo struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

// TableSchema defines the schema of a database table
type TableSchema struct {
	Columns []ColumnSchema `json:"columns,omitempty"`
}

// ColumnSchema defines the schema of a column in a database table
type ColumnSchema struct {
	Name         string `json:"name,omitempty"`
	DataType     string `json:"dataType,omitempty"`
	IsNullable   bool   `json:"isNullable,omitempty"`
	IsPrimaryKey bool   `json:"isPrimaryKey,omitempty"`
}
