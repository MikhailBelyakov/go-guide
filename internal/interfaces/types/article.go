package article

type RequestCreate struct {
	Title string `json:"title"`
}

type ResponseCreate struct {
	ID string `json:"id"`
}

type ResponseOne struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
