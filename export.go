package prosemirror

import (
	"bytes"
	"strings"
)

const tab = "  "

func renderContent(content *Content, config *Config, buf *bytes.Buffer, depth int, index int, parent *Content) {
	if content.Type == "list_item" {
		buf.WriteString(strings.Repeat(tab, depth-1))
		buf.WriteString(config.GetListNodeBefore(index, content, parent.Type))
	} else {
		buf.WriteString(config.GetNodeBefore(index, content))
	}
	if content.Text != "" {
		for i, mark := range content.Marks {
			buf.WriteString(config.GetMarkBefore(i, mark))
		}
		buf.WriteString(content.Text)
		for i := len(content.Marks) - 1; i >= 0; i-- {
			buf.WriteString(config.GetMarkAfter(i, content.Marks[i]))
		}
	} else if content.Content != nil {
		for i, c := range content.Content {
			if c.Type == "bullet_list" || c.Type == "ordered_list" {
				if depth == 0 {
					buf.WriteRune('\n')
				}
				renderContent(c, config, buf, depth+1, i, content)
				if depth == 0 {
					buf.WriteRune('\n')
				}
			} else {
				renderContent(c, config, buf, depth, i, content)
			}
		}
	}
	if content.Type == "list_item" {
		buf.WriteString(config.GetListNodeAfter(index, content, parent.Type))
	} else {
		buf.WriteString(config.GetNodeAfter(index, content))
	}
}

func Render(editorState *EditorState, config *Config) string {
	var buf bytes.Buffer
	renderContent(editorState.Doc, config, &buf, 0, 0, nil)
	return strings.TrimSuffix(buf.String(), "\n")
}
