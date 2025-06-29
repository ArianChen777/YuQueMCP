package tools

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/mcp"

	"yuque-mcp/internal/yuque"
)

// RegisterTools 注册所有语雀工具
func RegisterTools(server mcp.McpServer, client *yuque.Client) {
	registerGetDocTool(server, client)
	registerCreateDocTool(server, client)
	registerUpdateDocTool(server, client)
	registerSearchDocsTool(server, client)
	registerGetCurrentUserTool(server, client)
}

// 获取当前用户信息工具
func registerGetCurrentUserTool(server mcp.McpServer, client *yuque.Client) {
	tool := mcp.Tool{
		Name:        "yuque_get_current_user",
		Description: "获取当前语雀用户的详细信息",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{},
			Required:   []string{},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			user, err := client.GetCurrentUser()
			if err != nil {
				return nil, fmt.Errorf("获取用户信息失败: %w", err)
			}

			return map[string]any{
				"id":          user.ID,
				"login":       user.Login,
				"name":        user.Name,
				"description": user.Description,
				"avatar_url":  user.Avatar,
				"created_at":  user.CreatedAt,
				"updated_at":  user.UpdatedAt,
			}, nil
		},
	}
	server.RegisterTool(tool)
}

// 获取文档工具
func registerGetDocTool(server mcp.McpServer, client *yuque.Client) {
	tool := mcp.Tool{
		Name:        "yuque_get_document",
		Description: "根据知识库ID和文档ID获取语雀文档的详细内容",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"book_id": map[string]any{
					"type":        "string",
					"description": "知识库ID或路径（如：user/repo 或数字ID）",
				},
				"doc_id": map[string]any{
					"type":        "string",
					"description": "文档ID或slug",
				},
			},
			Required: []string{"book_id", "doc_id"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				BookID string `json:"book_id"`
				DocID  string `json:"doc_id"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("参数解析失败: %w", err)
			}

			doc, err := client.GetDoc(req.BookID, req.DocID)
			if err != nil {
				return nil, fmt.Errorf("获取文档失败: %w", err)
			}

			return map[string]any{
				"id":                 doc.ID,
				"title":              doc.Title,
				"slug":               doc.Slug,
				"body":               doc.Body,
				"body_html":          doc.BodyHTML,
				"format":             doc.Format,
				"public":             doc.Public,
				"status":             doc.Status,
				"likes_count":        doc.LikesCount,
				"comments_count":     doc.CommentsCount,
				"content_updated_at": doc.ContentUpdatedAt,
				"created_at":         doc.CreatedAt,
				"updated_at":         doc.UpdatedAt,
				"book_id":            doc.BookID,
				"user_id":            doc.UserID,
			}, nil
		},
	}
	server.RegisterTool(tool)
}

// 创建文档工具
func registerCreateDocTool(server mcp.McpServer, client *yuque.Client) {
	tool := mcp.Tool{
		Name:        "yuque_create_document",
		Description: "在指定的语雀知识库中创建新文档",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"book_id": map[string]any{
					"type":        "string",
					"description": "知识库ID或路径（如：user/repo 或数字ID）",
				},
				"title": map[string]any{
					"type":        "string",
					"description": "文档标题",
				},
				"body": map[string]any{
					"type":        "string",
					"description": "文档内容（支持Markdown格式）",
					"default":     "",
				},
				"format": map[string]any{
					"type":        "string",
					"description": "文档格式",
					"enum":        []string{"markdown", "lake"},
					"default":     "markdown",
				},
				"public": map[string]any{
					"type":        "integer",
					"description": "公开设置：0=私密，1=公开",
					"default":     0,
				},
			},
			Required: []string{"book_id", "title"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				BookID string `json:"book_id"`
				Title  string `json:"title"`
				Body   string `json:"body,optional"`
				Format string `json:"format,optional"`
				Public int    `json:"public,optional"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("参数解析失败: %w", err)
			}

			// 设置默认值
			if req.Format == "" {
				req.Format = "markdown"
			}

			createReq := yuque.CreateDocRequest{
				Title:  req.Title,
				Body:   req.Body,
				Format: req.Format,
				Public: req.Public,
			}

			doc, err := client.CreateDoc(req.BookID, createReq)
			if err != nil {
				return nil, fmt.Errorf("创建文档失败: %w", err)
			}

			return map[string]any{
				"id":         doc.ID,
				"title":      doc.Title,
				"slug":       doc.Slug,
				"body":       doc.Body,
				"format":     doc.Format,
				"public":     doc.Public,
				"created_at": doc.CreatedAt,
				"book_id":    doc.BookID,
				"message":    fmt.Sprintf("文档 '%s' 创建成功", doc.Title),
			}, nil
		},
	}
	server.RegisterTool(tool)
}

// 更新文档工具
func registerUpdateDocTool(server mcp.McpServer, client *yuque.Client) {
	tool := mcp.Tool{
		Name:        "yuque_update_document",
		Description: "更新指定的语雀文档内容",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"book_id": map[string]any{
					"type":        "string",
					"description": "知识库ID或路径（如：user/repo 或数字ID）",
				},
				"doc_id": map[string]any{
					"type":        "string",
					"description": "文档ID或slug",
				},
				"title": map[string]any{
					"type":        "string",
					"description": "文档标题（可选）",
				},
				"body": map[string]any{
					"type":        "string",
					"description": "文档内容（可选，支持Markdown格式）",
				},
				"format": map[string]any{
					"type":        "string",
					"description": "文档格式（可选）",
					"enum":        []string{"markdown", "lake"},
				},
				"public": map[string]any{
					"type":        "integer",
					"description": "公开设置：0=私密，1=公开（可选）",
				},
			},
			Required: []string{"book_id", "doc_id"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				BookID string `json:"book_id"`
				DocID  string `json:"doc_id"`
				Title  string `json:"title,optional"`
				Body   string `json:"body,optional"`
				Format string `json:"format,optional"`
				Public *int   `json:"public,optional"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("参数解析失败: %w", err)
			}

			updateReq := yuque.UpdateDocRequest{
				Title:  req.Title,
				Body:   req.Body,
				Format: req.Format,
			}

			if req.Public != nil {
				updateReq.Public = *req.Public
			}

			doc, err := client.UpdateDoc(req.BookID, req.DocID, updateReq)
			if err != nil {
				return nil, fmt.Errorf("更新文档失败: %w", err)
			}

			return map[string]any{
				"id":         doc.ID,
				"title":      doc.Title,
				"slug":       doc.Slug,
				"body":       doc.Body,
				"format":     doc.Format,
				"public":     doc.Public,
				"updated_at": doc.UpdatedAt,
				"book_id":    doc.BookID,
				"message":    fmt.Sprintf("文档 '%s' 更新成功", doc.Title),
			}, nil
		},
	}
	server.RegisterTool(tool)
}

// 搜索文档工具
func registerSearchDocsTool(server mcp.McpServer, client *yuque.Client) {
	tool := mcp.Tool{
		Name:        "yuque_search_documents",
		Description: "在语雀中搜索文档",
		InputSchema: mcp.InputSchema{
			Properties: map[string]any{
				"query": map[string]any{
					"type":        "string",
					"description": "搜索关键词",
				},
			},
			Required: []string{"query"},
		},
		Handler: func(ctx context.Context, params map[string]any) (any, error) {
			var req struct {
				Query string `json:"query"`
			}

			if err := mcp.ParseArguments(params, &req); err != nil {
				return nil, fmt.Errorf("参数解析失败: %w", err)
			}

			result, err := client.SearchDocs(req.Query)
			if err != nil {
				return nil, fmt.Errorf("搜索失败: %w", err)
			}

			// 格式化搜索结果
			documents := make([]map[string]any, 0, len(result.Hits))
			for _, hit := range result.Hits {
				doc := map[string]any{
					"id":       hit.Source.ID,
					"title":    hit.Source.Title,
					"slug":     hit.Source.Slug,
					"book_id":  hit.Source.BookID,
					"user_id":  hit.Source.UserID,
					"format":   hit.Source.Format,
					"created_at": hit.Source.CreatedAt,
					"updated_at": hit.Source.UpdatedAt,
				}

				// 添加高亮信息
				if len(hit.Highlight.Title) > 0 {
					doc["highlighted_title"] = hit.Highlight.Title[0]
				}
				if len(hit.Highlight.BodyHTML) > 0 {
					doc["highlighted_content"] = hit.Highlight.BodyHTML[0]
				}

				documents = append(documents, doc)
			}

			return map[string]any{
				"query":     result.Query,
				"total":     result.Total,
				"documents": documents,
				"message":   fmt.Sprintf("找到 %d 个相关文档", result.Total),
			}, nil
		},
	}
	server.RegisterTool(tool)
}