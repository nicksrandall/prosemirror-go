package prosemirror

type Option interface {
	RenderBefore(int, map[string]interface{}) string
	RenderAfter(int, map[string]interface{}) string
}

type Config struct {
	NodeRenderers map[string]Option
	MarkRenderers map[string]Option
}

const ListTypeKey = "__LIST_TYPE__"

func (c *Config) GetNodeBefore(i int, content *Content) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		return r.RenderBefore(i, content.Attrs)
	}
	return ""
}
func (c *Config) GetListNodeBefore(i int, content *Content, listType string) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		if content.Attrs == nil {
			content.Attrs = make(map[string]interface{})
		}
		content.Attrs[ListTypeKey] = listType
		return r.RenderBefore(i, content.Attrs)
	}
	return ""
}
func (c *Config) GetNodeAfter(i int, content *Content) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		return r.RenderAfter(i, content.Attrs)
	}
	return ""
}
func (c *Config) GetListNodeAfter(i int, content *Content, listType string) string {
	if r, ok := c.NodeRenderers[content.Type]; ok {
		if content.Attrs == nil {
			content.Attrs = make(map[string]interface{})
		}
		content.Attrs[ListTypeKey] = listType
		return r.RenderAfter(i, content.Attrs)
	}
	return ""
}
func (c *Config) GetMarkBefore(i int, mark *Mark) string {
	if r, ok := c.MarkRenderers[mark.Type]; ok {
		return r.RenderBefore(i, mark.Attrs)
	}
	return ""
}
func (c *Config) GetMarkAfter(i int, mark *Mark) string {
	if r, ok := c.MarkRenderers[mark.Type]; ok {
		return r.RenderAfter(i, mark.Attrs)
	}
	return ""
}

type SimpleOption struct {
	Before string
	After  string
}

func (o SimpleOption) RenderBefore(_ int, _ map[string]interface{}) string {
	return o.Before
}

func (o SimpleOption) RenderAfter(_ int, _ map[string]interface{}) string {
	return o.After
}

// Compile type check
var _ Option = (*SimpleOption)(nil)
