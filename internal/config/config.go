package config

import "github.com/zeromicro/go-zero/mcp"

type Config struct {
	mcp.McpConf
	YuQue YuQueConfig `json:",optional"`
}

type YuQueConfig struct {
	BaseURL string `json:",default=https://www.yuque.com/api/v2"`
	Token   string `json:",optional,env=YUQUE_TOKEN"`
}