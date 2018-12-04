package mobiledoc

type node struct {
	tagname, value string
	children       []*node
	attributes     map[string]string
}

func newNode(tagname, value string) *node {
	return &node{
		tagname:    tagname,
		value:      value,
		attributes: make(map[string]string),
	}
}

func (e *node) addChild(c *node) {
	e.children = append(e.children, c)
}

func (e *node) addAttribute(key, value string) {
	e.attributes[key] = value
}
