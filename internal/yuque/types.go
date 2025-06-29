package yuque

// User 表示语雀用户
type User struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Login       string `json:"login"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Avatar      string `json:"avatar_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Doc 表示文档基本信息
type Doc struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	BookID    int    `json:"book_id"`
	Book      *Book  `json:"book,omitempty"`
	UserID    int    `json:"user_id"`
	User      *User  `json:"user,omitempty"`
	Format    string `json:"format"`
	Body      string `json:"body"`
	BodyDraft string `json:"body_draft"`
	BodyHTML  string `json:"body_html"`
	Public    int    `json:"public"`
	Status    int    `json:"status"`
	LikesCount     int       `json:"likes_count"`
	CommentsCount  int       `json:"comments_count"`
	ContentUpdatedAt string  `json:"content_updated_at"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

// DocDetail 表示文档详细信息
type DocDetail struct {
	Doc
	Creator *User `json:"creator,omitempty"`
}

// Book 表示知识库
type Book struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	CreatorID   int    `json:"creator_id"`
	Public      int    `json:"public"`
	ItemsCount  int    `json:"items_count"`
	LikesCount  int    `json:"likes_count"`
	WatchesCount int   `json:"watches_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	User        *User  `json:"user,omitempty"`
	Creator     *User  `json:"creator,omitempty"`
}

// CreateDocRequest 创建文档请求
type CreateDocRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body,omitempty"`
	Format string `json:"format,omitempty"` // markdown, lake
	Public int    `json:"public,omitempty"` // 0: 私密, 1: 公开
}

// UpdateDocRequest 更新文档请求
type UpdateDocRequest struct {
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	Format string `json:"format,omitempty"`
	Public int    `json:"public,omitempty"`
}

// SearchResult 搜索结果
type SearchResult struct {
	Query string `json:"q"`
	Hits  []struct {
		Type      string `json:"type"`
		Source    Doc    `json:"_source"`
		Highlight struct {
			BodyHTML []string `json:"body_html,omitempty"`
			Title    []string `json:"title,omitempty"`
		} `json:"highlight,omitempty"`
	} `json:"hits"`
	Total int `json:"total"`
}