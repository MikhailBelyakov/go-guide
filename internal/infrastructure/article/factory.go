package article

func NewArticleModel(id, title string) *Entity {
	return &Entity{
		ID:    id,
		Title: title,
	}
}
