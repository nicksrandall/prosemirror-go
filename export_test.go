package prosemirror

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testJSON = `{"doc":{"type":"doc","content":[{"type":"heading","attrs":{"level":1},"content":[{"type":"text","text":"This is a header"}]},{"type":"paragraph","content":[{"type":"text","text":"This is a "},{"type":"text","marks":[{"type":"em"},{"type":"strong"}],"text":"bold, underlined"},{"type":"text","text":", message."}]},{"type":"ordered_list","content":[{"type":"list_item","content":[{"type":"paragraph","content":[{"type":"text","text":"this is a list"}]}]},{"type":"list_item","content":[{"type":"paragraph","content":[{"type":"text","text":"so is this"}]},{"type":"ordered_list","content":[{"type":"list_item","content":[{"type":"paragraph","content":[{"type":"text","text":"Nested list"}]}]}]}]}]},{"type":"paragraph"},{"type":"paragraph","content":[{"type":"text","text":"Another "},{"type":"variable","attrs":{"name":"Amount.BillingCount","value":"Amount_BillingCount","type":"string"}}]},{"type":"paragraph","content":[{"type":"text","marks":[{"type":"link","attrs":{"href":"http://www.domo.com","title":"here we go"}}],"text":"Testing"},{"type":"text","text":"."}]}]},"selection":{"type":"text","anchor":119,"head":119}}`

const expectedHTML = "<h1>This is a header</h1><p>This is a <em><strong>bold, underlined</strong></em>, message.</p>\n<ol><li><p>this is a list</p></li><li><p>so is this</p><ol>  <li><p>Nested list</p></li></ol></li></ol>\n<p></p><p>Another {Amount_BillingCount}</p><p><a href=\"http://www.domo.com\" title=\"here we go\">Testing</a>.</p>"
const expectedPlainText = "This is a header\nThis is a *bold, underlined*, message.\n\n1. this is a list\n2. so is this\n  1. Nested list\n\n\nAnother {Amount_BillingCount}\n[Testing](http://www.domo.com)."

func Test_Render_HTML(t *testing.T) {
	var editorState EditorState

	if err := json.Unmarshal([]byte(testJSON), &editorState); err != nil {
		t.Errorf("failed to Unmarshal JSON: %v", err)
	}

	config := NewHTMLConfig()
	html := Render(&editorState, config)

	assert.Equal(t, expectedHTML, html, "They should be equal")
}

func Test_Render_PlainText(t *testing.T) {
	var editorState EditorState

	if err := json.Unmarshal([]byte(testJSON), &editorState); err != nil {
		t.Errorf("failed to Unmarshal JSON: %v", err)
	}

	config := NewPlainTextConfig()
	text := Render(&editorState, config)

	assert.Equal(t, expectedPlainText, text, "They should be equal")
}
