package mobiledoc

type node struct {
	parent, firstChild, lastChild, prevSibling, nextSibling *node

	tagname, value string
	attributes     map[string]string
}

func newNode(tagname, value string) *node {
	return &node{
		tagname:    tagname,
		value:      value,
		attributes: make(map[string]string),
	}
}

// appendChild adds a node c as a child of n.
//
// It will panic if c already has a parent or siblings.
func (n *node) appendChild(c *node) {
	if c.parent != nil || c.prevSibling != nil || c.nextSibling != nil {
		panic("node: appendChild called for an attached child Node")
	}
	last := n.lastChild
	if last != nil {
		last.nextSibling = c
	} else {
		n.firstChild = c
	}
	n.lastChild = c
	c.parent = n
	c.prevSibling = last
}

// addAttribute adds an attribute key with value value.
func (n *node) addAttribute(key, value string) {
	n.attributes[key] = value
}
