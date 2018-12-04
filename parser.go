package mobiledoc

import (
	"encoding/json"

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
