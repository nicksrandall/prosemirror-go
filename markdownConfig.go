package prosemirror

import "strings"

func NewMarkdownConfig() *Config {
	nodeRenderers := map[string]Option{
		"paragraph":       SimpleOption{After: "\n"},
		"blockquote":      SimpleOption{Before: "> ", After: "\n"},
		"horizontal_rule": SimpleOption{Before: "---\n"},
		"heading":         markdownHeadingOption{},
		"hard_break":      SimpleOption{Before: "\n\n"},
		"ordered_list":    SimpleOption{After: "\n"},
		"bullet_list":     SimpleOption{After: "\n"},
		"list_item":       plainListItemOption{},
		"code_block":      SimpleOption{Before: "```\n", After: "```\n"},
	}

	markRenderers := map[string]Option{
		"link":   plainLinkOption{},
		"em":     SimpleOption{Before: "_", After: "_"},
		"strong": SimpleOption{Before: "**", After: "**"},
	}

	return &Config{nodeRenderers, markRenderers}
}

type markdownHeadingOption struct{}

func (o markdownHeadingOption) RenderBefore(_ int, attrs map[string]interface{}) string {
	return strings.Repeat("#", int(attrs["level"].(float64))) + " "
}

func (o markdownHeadingOption) RenderAfter(_ int, attrs map[string]interface{}) string {
	return ""
}
