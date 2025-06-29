# 贡献指南

感谢您对 YuQue MCP 服务器项目的关注！我们欢迎所有形式的贡献。

## 参与方式

### 🐛 报告Bug
- 使用 [Bug报告模板](https://github.com/ArianChen777/YuQueMCP/issues/new?template=bug_report.md) 创建issue
- 提供详细的复现步骤和环境信息
- 如可能，请提供相关的日志输出

### ✨ 功能建议
- 使用 [功能请求模板](https://github.com/ArianChen777/YuQueMCP/issues/new?template=feature_request.md) 创建issue
- 详细描述功能需求和使用场景
- 说明功能的优先级和重要性

### ❓ 问题咨询
- 使用 [问题咨询模板](https://github.com/ArianChen777/YuQueMCP/issues/new?template=question.md) 创建issue
- 查看现有的issue，避免重复提问
- 提供详细的环境和配置信息

### 💻 代码贡献

#### 开发环境准备
1. Fork 此仓库
2. 克隆您fork的仓库：
   ```bash
   git clone https://github.com/YOUR_USERNAME/YuQueMCP.git
   cd YuQueMCP
   ```
3. 安装依赖：
   ```bash
   go mod tidy
   ```
4. 设置环境变量：
   ```bash
   export YUQUE_TOKEN="your_yuque_token"
   ```

#### 开发流程
1. **创建分支**：
   ```bash
   git checkout -b feature/your-feature-name
   # 或
   git checkout -b fix/your-bug-fix
   ```

2. **编写代码**：
   - 遵循Go代码规范
   - 添加必要的注释
   - 确保代码可读性

3. **测试**：
   ```bash
   go test ./...
   go run cmd/server/main.go -f config.yaml
   ```

4. **提交代码**：
   ```bash
   git add .
   git commit -m "feat: add new feature" # 使用约定式提交格式
   ```

5. **推送并创建PR**：
   ```bash
   git push origin feature/your-feature-name
   ```
   然后在GitHub上创建Pull Request

## 代码规范

### 提交信息格式
使用 [约定式提交](https://www.conventionalcommits.org/zh-hans/) 格式：

```
<类型>[可选的作用域]: <描述>

[可选的正文]

[可选的脚注]
```

**类型**：
- `feat`: 新功能
- `fix`: 修复bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 重构
- `perf`: 性能优化
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

**示例**：
```
feat(client): add document search functionality
fix(auth): handle token expiration correctly
docs: update API documentation
```

### Go代码规范
- 遵循 [Go代码审查注释](https://go.dev/wiki/CodeReviewComments)
- 使用 `gofmt` 格式化代码
- 使用有意义的变量和函数名
- 添加必要的错误处理
- 为公共函数添加文档注释

### 项目结构
```
├── cmd/server/          # 服务器入口
├── internal/
│   ├── config/         # 配置相关
│   ├── server/         # MCP服务器实现
│   │   └── tools/      # 工具实现
│   └── yuque/          # 语雀API客户端
├── .github/            # GitHub模板
├── config.yaml         # 配置文件
└── README.md          # 项目说明
```

## 开发建议

### 添加新功能
1. 首先创建feature request issue讨论需求
2. 确保功能符合MCP协议规范
3. 考虑向后兼容性
4. 添加相应的测试
5. 更新文档

### 修复Bug
1. 首先重现bug并创建测试用例
2. 修复问题
3. 确保测试通过
4. 验证不会引入新问题

### 性能优化
1. 测量当前性能
2. 实施优化
3. 对比优化前后的性能数据
4. 确保功能正确性不受影响

## 语雀API相关

### API文档
- [语雀开放API文档](https://app.swaggerhub.com/apis-docs/Jeff-Tian/yuque-open_api/2.0.1)

### 测试注意事项
- 使用测试用的语雀账号
- 不要在测试中删除重要数据
- 注意API频率限制
- 保护Token等敏感信息

## MCP协议相关

### 参考资源
- [MCP官方文档](https://modelcontextprotocol.io/docs)
- [go-zero MCP组件](https://github.com/zeromicro/go-zero/tree/master/mcp)

### 开发注意事项
- 遵循MCP协议规范
- 正确处理JSON-RPC格式
- 实现完整的错误处理
- 支持SSE通信

## 社区行为准则

### 我们期望的行为
- 使用友好和包容的语言
- 尊重不同的观点和经验
- 优雅地接受建设性批评
- 关注对社区最有利的事情
- 对其他社区成员表示同理心

### 不当行为
- 使用性化的语言或图像
- 发表侮辱/诽谤性评论，人身攻击
- 公开或私下骚扰
- 未经明确许可发布他人私人信息
- 其他在专业环境中被认为不当的行为

## 获得帮助

如果您在贡献过程中遇到问题：

1. 查看现有的 [Issues](https://github.com/ArianChen777/YuQueMCP/issues)
2. 创建新的 [问题咨询](https://github.com/ArianChen777/YuQueMCP/issues/new?template=question.md)
3. 参考项目文档和相关API文档

感谢您的贡献！🎉