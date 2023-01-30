package schema

type BookCreateRequest struct {
	Title      string   `json:"title" validate:"required"`
	Author     string   `json:"author" validate:"required"`
	Uri        string   `json:"uri" validate:"required"`
	Categories []string `json:"categories"`
}
