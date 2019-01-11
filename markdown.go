package mobiledoc

import (
	"fmt"
	"io"
	"strings"
)

func (n *node) renderEnd(w io.Writer) error {
	var err error
	switch strings.ToLower(n.tagname) {
	case BOLD, STRONG:
		_, err = fmt.Fprint(w, "**")
	case ITALIC, EMPHASIS:
		_, err = fmt.Fprint(w, "*")
	case ANCHOR:
		if _, err = fmt.Fprint(w, "]"); err != nil {
			return err
		}
		if href, ok := n.attributes["href"]; ok {
			_, err = fmt.Fprintf(w, "(%s)", href)
		}
	case IMAGE:
		if _, err = fmt.Fprint(w, "]"); err != nil {
			return err
		}
		if src, ok := n.attributes["src"]; ok {
			_, err = fmt.Fprintf(w, "(%s)", src)
		}
	case LISTITEM, ORDEREDLIST, UNORDEREDLIST:
		_, err = fmt.Fprint(w, "\n")
	case H1, H2, H3, H4, PARAGRAPH, BLOCKQUOTE:
		_, err = fmt.Fprint(w, "\n\n")
	}

	return err
}

func (n *node) renderListItemStart(w io.Writer) error {
	var err error
	if pos, ok := n.attributes["position"]; ok {
		_, err = fmt.Fprintf(w, "%s. ", pos)
	} else {
		_, err = fmt.Fprint(w, "* ")
	}
	return err
}

func (n *node) renderStart(w io.Writer) error {
	var err error
	switch strings.ToLower(n.tagname) {
	case BOLD, STRONG:
		_, err = fmt.Fprint(w, "**")
	case ITALIC, EMPHASIS:
		_, err = fmt.Fprint(w, "*")
	case H1:
		_, err = fmt.Fprint(w, "# ")
	case H2:
		_, err = fmt.Fprint(w, "## ")
	case H3:
		_, err = fmt.Fprint(w, "### ")
	case H4:
		_, err = fmt.Fprint(w, "#### ")
	case ANCHOR:
		_, err = fmt.Fprint(w, "[")
	case IMAGE:
		_, err = fmt.Fprint(w, "![")
	case LISTITEM:
		err = n.renderListItemStart(w)
	case BLOCKQUOTE:
		_, err = fmt.Fprint(w, "> ")
	}

	return err
}

func (n *node) renderMarkdown(w io.Writer) error {
	if err := n.renderStart(w); err != nil {
		return err
	}

	if n.value != "" {
		// Has a value, so render
		if _, err := fmt.Fprint(w, n.value); err != nil {
			return err
		}
	} else {
		// No value, so render children
		for _, c := range n.children {
			err := c.renderMarkdown(w)
			if err != nil {
				return err
			}
		}
	}

	return n.renderEnd(w)
}
