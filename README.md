# 语雀 MCP 服务器

基于 go-zero MCP 组件开发的语雀 Model Context Protocol 服务器，为 AI 助手提供语雀知识库操作能力。

## 功能特性

- 🔍 **获取文档**: 根据知识库ID和文档ID获取文档详细内容
- 📝 **创建文档**: 在指定知识库中创建新文档  
- ✏️ **更新文档**: 修改文档标题、内容等信息
- 🔎 **搜索文档**: 在语雀中搜索相关文档
- 👤 **用户信息**: 获取当前用户详细信息

## 快速开始

### 1. 设置环境变量

```bash
export YUQUE_TOKEN="your_yuque_token_here"
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 启动服务器

```bash
go run cmd/server/main.go -f config.yaml
```

服务器将在 `http://localhost:8080` 启动。

## 可用工具

### yuque_get_current_user
获取当前语雀用户信息。

### yuque_get_document
获取指定文档内容。
- `book_id`: 知识库ID或路径
- `doc_id`: 文档ID或slug

### yuque_create_document  
创建新文档。
- `book_id`: 知识库ID或路径
- `title`: 文档标题
- `body`: 文档内容（可选）
- `format`: 文档格式，默认markdown
- `public`: 公开设置，默认0（私密）

### yuque_update_document
更新文档内容。
- `book_id`: 知识库ID或路径  
- `doc_id`: 文档ID或slug
- `title`: 新标题（可选）
- `body`: 新内容（可选）
- `format`: 文档格式（可选）
- `public`: 公开设置（可选）

### yuque_search_documents
搜索文档。
- `query`: 搜索关键词

## 配置说明

配置文件 `config.yaml`:

```yaml
name: yuque-mcp-server
host: localhost
port: 8080
mcp:
  name: yuque-mcp-server
  messageTimeout: 30s
  cors:
    - http://localhost:3000

yuque:
  baseUrl: "https://www.yuque.com/api/v2"
  # token 通过环境变量 YUQUE_TOKEN 设置
```

## Cursor IDE 集成配置

在 Cursor IDE 中使用此 MCP 服务器，需要在设置中添加以下配置：

```json
{
  "mcpServers": {
    "yuque": {
      "command": "go",
      "args": ["run", "cmd/server/main.go", "-f", "config.yaml"],
      "cwd": "/path/to/YuQueMCP",
      "env": {
        "YUQUE_TOKEN": "your_yuque_token_here"
      }
    }
  }
}
```

配置步骤：
1. 打开 Cursor IDE 设置
2. 搜索 "MCP" 或 "Model Context Protocol"
3. 添加上述 JSON 配置
4. 将 `cwd` 路径修改为项目实际路径
5. 将 `YUQUE_TOKEN` 替换为你的语雀 API Token
6. 重启 Cursor IDE

## 项目结构

```
├── cmd/server/          # 服务器启动入口
├── internal/
│   ├── config/         # 配置相关
│   ├── server/         # MCP服务器逻辑
│   │   └── tools/      # 工具实现
│   └── yuque/          # 语雀API客户端
├── config.yaml         # 配置文件
└── go.mod             # Go模块文件
```