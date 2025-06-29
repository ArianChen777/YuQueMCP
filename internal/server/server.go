package server

import (
	"github.com/zeromicro/go-zero/mcp"

	"yuque-mcp/internal/config"
	"yuque-mcp/internal/server/tools"
	"yuque-mcp/internal/yuque"
)

// RegisterYuQueHandlers 注册语雀相关的工具、提示和资源
func RegisterYuQueHandlers(server mcp.McpServer, cfg config.YuQueConfig) {
	// 创建语雀客户端
	yuqueClient := yuque.NewClient(cfg)

	// 注册工具
	tools.RegisterTools(server, yuqueClient)

	// TODO: 注册提示和资源
	// prompts.RegisterPrompts(server, yuqueClient)
	// resources.RegisterResources(server, yuqueClient)
}