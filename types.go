package prosemirror

// EditorState - http://prosemirror.net/docs/ref/#state.Editor_State
type EditorState struct {
	Doc       *Content   `json:"doc"`
	Selection *Selection `json:"selection"`
}

// Content - http://prosemirror.net/docs/ref/#model.Node
type Content struct {
	Attrs   map[string]interface{} `json:"attrs,omitempty"`
	Content []*Content             `json:"content,omitempty"`
	Marks   []*Mark                `json:"marks,omitempty"`
	Text    string                 `json:"text,omitempty"`
	Type    string                 `json:"type"`
}

// Selection - http://prosemirror.net/docs/ref/#state.Selection
type Selection struct {
	Anchor int    `json:"anchor"`
	Head   int    `json:"head"`
	Type   string `json:"type"`
}

// Mark - http://prosemirror.net/docs/ref/#model.Mark
type Mark struct {
	Attrs map[string]interface{} `json:"attrs,omitempty"`
	Type  string                 `json:"type"`
}
