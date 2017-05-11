package prosemirror

type EditorState struct {
	Doc       *Content   `json:"doc"`
	Selection *Selection `json:"selection"`
}

type Content struct {
	Attrs   map[string]interface{} `json:"attrs,omitempty"`
	Content []*Content             `json:"content,omitempty"`
	Marks   []*Mark                `json:"marks,omitempty"`
	Text    string                 `json:"text,omitempty"`
	Type    string                 `json:"type"`
}

type Selection struct {
	Anchor int    `json:"anchor"`
	Head   int    `json:"head"`
	Type   string `json:"type"`
}

type Mark struct {
	Attrs map[string]interface{} `json:"attrs,omitempty"`
	Type  string                 `json:"type"`
}
