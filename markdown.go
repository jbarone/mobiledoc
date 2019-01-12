package mobiledoc

import (
	"fmt"
	"io"
	"strings"
)

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
		_, err = fmt.Fprint(w, "_")
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

func (n *node) renderEnd(w io.Writer) error {
	var err error
	switch strings.ToLower(n.tagname) {
	case BOLD, STRONG:
		_, err = fmt.Fprint(w, "**")
	case ITALIC, EMPHASIS:
		_, err = fmt.Fprint(w, "_")
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

func (n *node) rednerSpace(w io.Writer) error {
	var err error

	switch strings.ToLower(n.tagname) {
	case LISTITEM, ORDEREDLIST, UNORDEREDLIST, H1, H2, H3, H4, PARAGRAPH, BLOCKQUOTE:
		// do nothing
	default:
		_, err = fmt.Fprint(w, " ")
	}

	return err
}

func (n *node) renderContent(w io.Writer) error {
	var err error
	if n.value != "" {
		_, err = fmt.Fprint(w, strings.TrimSpace(n.value))
		return err
	}
	for c := n.firstChild; c != nil; c = c.nextSibling {
		if err = c.renderMarkdown(w); err != nil {
			return err
		}
		if c.nextSibling != nil {
			if err = c.rednerSpace(w); err != nil {
				return err
			}
		}
	}
	return err
}

func (n *node) renderMarkdown(w io.Writer) error {
	var err error
	if err = n.renderStart(w); err != nil {
		return err
	}

	if err = n.renderContent(w); err != nil {
		return err
	}

	err = n.renderEnd(w)
	return err
}
