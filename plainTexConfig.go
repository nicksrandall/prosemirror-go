package prosemirror

import "fmt"

// NewPlaintextConfig creates a config configured to render plain text
func NewPlainTextConfig() *Config {
	nodeRenderers := map[string]Option{
		"paragraph":       SimpleOption{After: "\n"},
		"blockquote":      SimpleOption{Before: "| ", After: "\n"},
		"horizontal_rule": SimpleOption{Before: "---\n"},
		"heading":         SimpleOption{After: "\n"},
		"hard_break":      SimpleOption{Before: "\n"},
		"list_item":       plainListItemOption{},
		"variable":        variableOption{},
	}

	markRenderers := map[string]Option{
		"link":   plainLinkOption{},
		"em":     SimpleOption{Before: "", After: ""},
		"strong": SimpleOption{Before: "*", After: "*"},
		"code":   SimpleOption{Before: "", After: ""},
	}

	return &Config{nodeRenderers, markRenderers}
}

type plainLinkOption struct{}

func (o plainLinkOption) RenderBefore(_ int, attrs map[string]interface{}) string {
	return "["
}
func (o plainLinkOption) RenderAfter(_ int, attrs map[string]interface{}) string {
	return fmt.Sprintf("](%s)", attrs["href"])
}

type plainListItemOption struct{}

func (o plainListItemOption) RenderBefore(i int, attrs map[string]interface{}) string {
	if t, ok := attrs[ListTypeKey]; ok && t == "ordered_list" {
		return fmt.Sprintf("%d. ", i+1)
	}
	return "* "
}
func (o plainListItemOption) RenderAfter(_ int, _ map[string]interface{}) string {
	return ""
}

// Compile type check
var (
	_ Option = (*plainLinkOption)(nil)
)
