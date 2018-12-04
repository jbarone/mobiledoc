package mobiledoc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type doc struct {
	markups []markup
	atoms   []atom
	cards   []card
}

// nolint [: gocyclo]
func parseV03(mdmap map[string]json.RawMessage) (*node, error) {
	root := newNode(DIV, "")
	var d doc

	if markups, ok := mdmap["markups"]; ok {
		err := json.Unmarshal(markups, &d.markups)
		if err != nil {
			return root, err
		}
	}

	if atoms, ok := mdmap["atoms"]; ok {
		err := json.Unmarshal(atoms, &d.atoms)
		if err != nil {
			return root, err
		}
	}

	if cards, ok := mdmap["cards"]; ok {
		err := json.Unmarshal(cards, &d.cards)
		if err != nil {
			return root, err
		}
	}

	sections, ok := mdmap["sections"]
	if !ok {
		return root, errors.New("invalid mobiledoc: sections missing")
	}

	var rawSections [][]json.RawMessage
	err := json.Unmarshal(sections, &rawSections)
	if err != nil {
		return root, err
	}

	for _, s := range rawSections {
		var t int
		err = json.Unmarshal(s[0], &t)
		if err != nil {
			return root, err
		}
		switch t {
		case sectionImage:
			var url string
			err = json.Unmarshal(s[1], &url)
			if err != nil {
				return root, err
			}
			n := newNode(IMAGE, "")
			n.addAttribute("src", url)
			root.addChild(n)

		case sectionList:
			var tag string
			err = json.Unmarshal(s[1], &tag)
			if err != nil {
				return root, err
			}
			switch strings.ToLower(tag) {
			case UNORDEREDLIST, ORDEREDLIST:
				// do nothing - valid
			default:
			}

			n := newNode(tag, "")

			var items [][]json.RawMessage
			err = json.Unmarshal(s[2], &items)
			if err != nil {
				return root, err
			}

			for pos, markers := range items {
				li := newNode(LISTITEM, "")
				if strings.ToLower(tag) == ORDEREDLIST {
					li.addAttribute("position", strconv.Itoa(pos+1))
				}

				err = addMarkersToNode(li, markers, d)
				if err != nil {
					return root, err
				}

				n.addChild(li)
			}

			root.addChild(n)

		case sectionMarkup:
			var tag string
			err = json.Unmarshal(s[1], &tag)
			if err != nil {
				return root, err
			}
			n := newNode(tag, "")

			var markers []json.RawMessage
			err = json.Unmarshal(s[2], &markers)
			if err != nil {
				return root, err
			}

			err = addMarkersToNode(n, markers, d)
			if err != nil {
				return root, err
			}

			root.addChild(n)

		case sectionCard:
			var cardIndex int
			err = json.Unmarshal(s[1], &cardIndex)
			if err != nil {
				return root, err
			}
			card := d.cards[cardIndex]
			n, err := card.Render("markdown")
			if err != nil {
				return root, err
			}
			root.addChild(n)
		}
	}

	return root, nil
}

// nolint [: gocyclo]
func addMarkersToNode(n *node, markers []json.RawMessage, d doc) error {
	nodes := []*node{n}
	for _, m := range markers {
		var mark marker
		err := json.Unmarshal(m, &mark)
		if err != nil {
			return err
		}

		for _, o := range mark.openIndexes {
			if o > len(d.markups) {
				return fmt.Errorf("unknown markup %d", o)
			}
			nodeMarkup := d.markups[o]
			switch strings.ToLower(nodeMarkup.tagName) {
			case BOLD, ITALIC, STRONG, EMPHASIS, ANCHOR, UNDERLINE,
				SUBSCRIPT, SUPERSCRIPT, STRIKETHROUGH:
				// do nothing - valid tag
			default:
				mark.closeCount--
				continue
			}
			mn := nodeMarkup.createNode()
			n.addChild(mn)
			nodes = append(nodes, mn)
			n = mn
		}

		switch mark.markerType {
		case markerMarkup:
			n.addChild(newNode(TEXT, mark.value.(string)))
		case markerAtom:
			atom := d.atoms[int(mark.value.(float64))]
			a, err := atom.Render("markdown")
			if err != nil {
				return err
			}
			n.addChild(a)
		}

		for i := 0; i < mark.closeCount; i++ {
			nodes = nodes[:len(nodes)-1]
			n = nodes[len(nodes)-1]
		}
	}

	return nil
}
