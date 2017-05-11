# prosemirror-go
[![GoDoc](https://godoc.org/github.com/nicksrandall/prosemirror-go?status.svg)](https://godoc.org/github.com/nicksrandall/prosemirror-go)
[![Build Status](https://travis-ci.org/nicksrandall/prosemirror-go.svg?branch=master)](https://travis-ci.org/nicksrandall/prosemirror-go)
This will export ProseMirror content state to HTML, Plain Text, or Markdown

## Usage

```go
func Export(rawContentState []byte) (string, error) {
  state := prosemirror.ContentState{}

  if err := json.Unmarshall(rawContentState, &state); err != nil {
    return "", err
  }

  config := prosemirror.NewHTMLConfig() // Export HTML
  // config := prosemirror.NewMarkdownConfig() // Export Markdown
  // config := prosemirror.NewPlainTextConfig() // Export Plain Text

  html := prosemirror.Render(&state, config)
  return html, nil
}

```
