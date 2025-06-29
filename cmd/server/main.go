package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/mcp"

	"yuque-mcp/internal/config"
	"yuque-mcp/internal/server"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 禁用统计日志
	logx.DisableStat()

	// 创建 MCP 服务器
	mcpServer := mcp.NewMcpServer(c.McpConf)
	defer mcpServer.Stop()

	// 注册语雀相关的工具、提示和资源
	server.RegisterYuQueHandlers(mcpServer, c.YuQue)

	fmt.Printf("Starting YuQue MCP Server on %s:%d\n", c.Host, c.Port)
	mcpServer.Start()
}
