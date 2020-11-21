package article

import (
	"goblog/pkg/route"
	"goblog/pkg/types"
)

type Article struct {
	ID    int64
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Int64ToString(a.ID))
}