package api

import (
	"errors"
)

// ClientFactory 客户端工厂接口
type ClientFactory interface {
	CreateClient(info ConnectionInfo) (Client, error)
}

// ClientFactories 客户端工厂注册表
var ClientFactories = make(map[DataSourceType]ClientFactory)

// RegisterClientFactory 注册客户端工厂
func RegisterClientFactory(dataSourceType DataSourceType, factory ClientFactory) {
	ClientFactories[dataSourceType] = factory
}

// NewClient 创建新的客户端
func NewClient(dataSourceType DataSourceType, info ConnectionInfo) (Client, error) {
	factory, ok := ClientFactories[dataSourceType]
	if !ok {
		return nil, errors.New("unsupported data source type")
	}
	return factory.CreateClient(info)
}
