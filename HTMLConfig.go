package prosemirror

import "fmt"

func NewHTMLConfig() *Config {
	nodeRenderers := map[string]Option{
		"paragraph":       SimpleOption{Before: "<p>", After: "</p>"},
		"blockquote":      SimpleOption{Before: "<blockquote>", After: "</blockquote>"},
		"horizontal_rule": SimpleOption{Before: "<hr>"},
		"heading":         headingOption{},
		"code_block":      SimpleOption{Before: "<pre>", After: "</pre>"},
		"text":            SimpleOption{Before: "", After: ""},
		"image":           SimpleOption{Before: "", After: ""},
		"hard_break":      SimpleOption{Before: "<br>"},
		"ordered_list":    SimpleOption{Before: "<ol>", After: "</ol>"},
		"bullet_list":     SimpleOption{Before: "<ul>", After: "</ul>"},
		"list_item":       SimpleOption{Before: "<li>", After: "</li>"},
	}

	markRenderers := map[string]Option{
		"link":   linkOption{},
		"em":     SimpleOption{Before: "<em>", After: "</em>"},
		"strong": SimpleOption{Before: "<strong>", After: "</strong>"},
		"code":   SimpleOption{Before: "<code>", After: "</code>"},
	}

	return &Config{nodeRenderers, markRenderers}
}

type headingOption struct{}

func (o headingOption) RenderBefore(_ int, attrs map[string]interface{}) string {
	return fmt.Sprintf("<h%d>", int(attrs["level"].(float64)))
}

func (o headingOption) RenderAfter(_ int, attrs map[string]interface{}) string {
	return fmt.Sprintf("</h%d>", int(attrs["level"].(float64)))
}

type linkOption struct{}

func (o linkOption) RenderBefore(_ int, attrs map[string]interface{}) string {
	return fmt.Sprintf("<a href=\"%s\" title=\"%s\">", attrs["href"], attrs["title"])
}
func (o linkOption) RenderAfter(_ int, attrs map[string]interface{}) string {
	return "</a>"
}

// Compile type check
var (
	_ Option = (*headingOption)(nil)
	_ Option = (*linkOption)(nil)
)
