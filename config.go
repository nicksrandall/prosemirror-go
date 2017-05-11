package prosemirror

// Option configures what to render before and after a specific content type
type Option interface {
	RenderBefore(int, map[string]interface{}) string
	RenderAfter(int, map[string]interface{}) string
}

// Config is passed to the render function and configures how to render the various content types
type Config struct {
	NodeRenderers map[string]Option
	MarkRenderers map[string]Option
}

// ListTypeKey is used to get the list type off of the attrs when rendering a list
const ListTypeKey = "__LIST_TYPE__"

func (c *Config) getNodeBefore(i int, content *Content) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		return r.RenderBefore(i, content.Attrs)
	}
	return ""
}
func (c *Config) getListNodeBefore(i int, content *Content, listType string) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		if content.Attrs == nil {
			content.Attrs = make(map[string]interface{})
		}
		content.Attrs[ListTypeKey] = listType
		return r.RenderBefore(i, content.Attrs)
	}
	return ""
}
func (c *Config) getNodeAfter(i int, content *Content) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		return r.RenderAfter(i, content.Attrs)
	}
	return ""
}
func (c *Config) getListNodeAfter(i int, content *Content, listType string) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		if content.Attrs == nil {
			content.Attrs = make(map[string]interface{})
		}
		content.Attrs[ListTypeKey] = listType
		return r.RenderAfter(i, content.Attrs)
	}
	return ""
}
func (c *Config) getMarkBefore(i int, mark *Mark) string {
	if r, ok := c.MarkRenderers[mark.Type]; ok {
		return r.RenderBefore(i, mark.Attrs)
	}
	return ""
}
func (c *Config) getMarkAfter(i int, mark *Mark) string {
	if r, ok := c.MarkRenderers[mark.Type]; ok {
		return r.RenderAfter(i, mark.Attrs)
	}
	return ""
}

// SimpleOption implements Option. Can be used when Before/After strings are constant.
type SimpleOption struct {
	Before string
	After  string
}

// RenderBefore returns the Before string
func (o SimpleOption) RenderBefore(_ int, _ map[string]interface{}) string {
	return o.Before
}

// RenderBefore returns the After string
func (o SimpleOption) RenderAfter(_ int, _ map[string]interface{}) string {
	return o.After
}

// Compile type check
var _ Option = (*SimpleOption)(nil)
