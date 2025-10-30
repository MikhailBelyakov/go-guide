package article

type Status string
type Type string

type Article struct {
	ID      string
	Title   string
	Content string
	Status  Status
	Type    Type
}
