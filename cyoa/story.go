package cyoa

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"story,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"arc,omitempty"`
}
