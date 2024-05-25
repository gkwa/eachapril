package example2

type MarkdownFile struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Path     string `json:"path"`
}
