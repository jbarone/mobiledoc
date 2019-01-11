package mobiledoc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Tag Types
const (
	BOLD          = "b"
	STRONG        = "strong"
	ITALIC        = "i"
	EMPHASIS      = "em"
	H1            = "h1"
	H2            = "h2"
	H3            = "h3"
	H4            = "h4"
	ANCHOR        = "a"
	IMAGE         = "img"
	LISTITEM      = "li"
	BLOCKQUOTE    = "blockquote"
	PARAGRAPH     = "p"
	DIV           = "div"
	ORDEREDLIST   = "ol"
	UNORDEREDLIST = "ul"
	UNDERLINE     = "u"
	SUBSCRIPT     = "sub"
	SUPERSCRIPT   = "sup"
	STRIKETHROUGH = "s"
	TEXT          = ""
)

type markup struct {
	tagName    string
	attributes map[string]string
}

func (m *markup) UnmarshalJSON(b []byte) error {
	var tmp []json.RawMessage
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal markup")
	}

	if len(tmp) == 0 {
		return errors.New("markup too short")
	}

	var tag string
	err = json.Unmarshal(tmp[0], &tag)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal markup tag name")
	}
	m.tagName = tag

	if len(tmp) == 1 {
		return nil
	}

	var attributes []string
	err = json.Unmarshal(tmp[1], &attributes)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal markup attributes")
	}
	if len(attributes)%2 != 0 {
		return errors.New("markup attributes must be in pairs")
	}

	m.attributes = make(map[string]string)
	for i := 0; i < len(attributes); i = i + 2 {
		m.attributes[attributes[i]] = attributes[i+1]
	}
	return nil
}

func (m *markup) createNode() *node {
	n := newNode(m.tagName, "")
	for k, v := range m.attributes {
		n.addAttribute(k, v)
	}
	return n
}

type marker struct {
	markerType  int
	openIndexes []int
	closeCount  int
	value       interface{}
}

func (m *marker) UnmarshalJSON(b []byte) error {
	var mark []json.RawMessage
	err := json.Unmarshal(b, &mark)
	if err != nil {
		return err
	}
	err = json.Unmarshal(mark[0], &m.markerType)
	if err != nil {
		return err
	}
	err = json.Unmarshal(mark[1], &m.openIndexes)
	if err != nil {
		return err
	}
	err = json.Unmarshal(mark[2], &m.closeCount)
	if err != nil {
		return err
	}
	err = json.Unmarshal(mark[3], &m.value)
	if err != nil {
		return err
	}
	return nil
}

func (md Mobiledoc) openMarker(n *node, m marker) (*node, []*node, error) {
	var nodes []*node
	for _, o := range m.openIndexes {
		if o > len(md.doc.markups) {
			return nil, nil, fmt.Errorf("unknown markup %d", o)
		}

		nodeMarkup := md.doc.markups[o]

		switch strings.ToLower(nodeMarkup.tagName) {
		case BOLD, ITALIC, STRONG, EMPHASIS, ANCHOR, UNDERLINE,
			SUBSCRIPT, SUPERSCRIPT, STRIKETHROUGH:
			// do nothing - valid tag
		default:
			m.closeCount--
			continue
		}

		node := nodeMarkup.createNode()
		n.addChild(node)
		nodes = append(nodes, node)
		n = node
	}

	return n, nodes, nil
}

func (md Mobiledoc) closeMarker(
	n *node, nodes []*node, m marker,
) (*node, []*node) {
	for i := 0; i < m.closeCount; i++ {
		nodes = nodes[:len(nodes)-1]
		n = nodes[len(nodes)-1]
	}
	return n, nodes
}
