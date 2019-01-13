package mobiledoc

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type doc struct {
	markups []markup
	atoms   []atom
	cards   []card
}

func parseDoc(mdmap map[string]json.RawMessage) (doc, error) {
	var d doc

	if markups, ok := mdmap["markups"]; ok {
		err := json.Unmarshal(markups, &d.markups)
		if err != nil {
			return doc{}, err
		}
	}

	if atoms, ok := mdmap["atoms"]; ok {
		err := json.Unmarshal(atoms, &d.atoms)
		if err != nil {
			return doc{}, err
		}
	}

	if cards, ok := mdmap["cards"]; ok {
		err := json.Unmarshal(cards, &d.cards)
		if err != nil {
			return doc{}, err
		}
	}
	return d, nil
}

func (md *Mobiledoc) parseSectionImage(root *node, s []json.RawMessage) error {
	var url string
	err := json.Unmarshal(s[1], &url)
	if err != nil {
		return err
	}
	n := newNode(IMAGE, "")
	n.addAttribute("src", url)
	root.appendChild(n)
	return nil
}

func (md *Mobiledoc) parseSectionList(root *node, s []json.RawMessage) error {
	var tag string
	err := json.Unmarshal(s[1], &tag)
	if err != nil {
		return err
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
		return err
	}

	for pos, markers := range items {
		li := newNode(LISTITEM, "")
		if strings.ToLower(tag) == ORDEREDLIST {
			li.addAttribute("position", strconv.Itoa(pos+1))
		}

		err = md.addMarkersToNode(li, markers)
		if err != nil {
			return err
		}

		n.appendChild(li)
	}

	root.appendChild(n)
	return nil
}

func (md *Mobiledoc) parseSectionMarkup(root *node, s []json.RawMessage) error {
	var tag string
	err := json.Unmarshal(s[1], &tag)
	if err != nil {
		return err
	}
	n := newNode(tag, "")

	var markers []json.RawMessage
	err = json.Unmarshal(s[2], &markers)
	if err != nil {
		return err
	}

	err = md.addMarkersToNode(n, markers)
	if err != nil {
		return err
	}

	root.appendChild(n)
	return err
}

func (md *Mobiledoc) parseSectionCard(root *node, s []json.RawMessage) error {
	var cardIndex int
	err := json.Unmarshal(s[1], &cardIndex)
	if err != nil {
		return err
	}
	card := md.doc.cards[cardIndex]
	n, err := md.renderCard(&card)
	if err != nil {
		return err
	}
	root.appendChild(n)
	return nil
}

func (md *Mobiledoc) parseSection(root *node, s []json.RawMessage) error {
	var t int
	err := json.Unmarshal(s[0], &t)
	if err != nil {
		return err
	}

	switch t {
	case sectionImage:
		if err = md.parseSectionImage(root, s); err != nil {
			return err
		}

	case sectionList:
		if err = md.parseSectionList(root, s); err != nil {
			return err
		}

	case sectionMarkup:
		if err = md.parseSectionMarkup(root, s); err != nil {
			return err
		}

	case sectionCard:
		if err = md.parseSectionCard(root, s); err != nil {
			return err
		}
	}
	return nil
}

func (md *Mobiledoc) parseV03(
	mdmap map[string]json.RawMessage,
) (*node, error) {
	root := newNode("root", "")

	d, err := parseDoc(mdmap)
	if err != nil {
		return root, err
	}
	md.doc = d

	sections, ok := mdmap["sections"]
	if !ok {
		return root, errors.New("invalid mobiledoc: sections missing")
	}

	var rawSections [][]json.RawMessage
	err = json.Unmarshal(sections, &rawSections)
	if err != nil {
		return root, err
	}

	for _, s := range rawSections {
		if err = md.parseSection(root, s); err != nil {
			return root, err
		}
	}

	return root, nil
}

func (md *Mobiledoc) addMarkersToNode(
	n *node, markers []json.RawMessage,
) error {
	nodes := []*node{n}
	for _, m := range markers {
		var mark marker
		err := json.Unmarshal(m, &mark)
		if err != nil {
			return err
		}

		node, openNodes, m, err := md.openMarker(n, mark)
		if err != nil {
			return err
		}
		nodes = append(nodes, openNodes...)
		n = node
		mark = m

		switch mark.markerType {
		case markerMarkup:
			n.appendChild(newNode(TEXT, mark.value.(string)))
		case markerAtom:
			atom := md.doc.atoms[int(mark.value.(float64))]
			a, err := md.renderAtom(&atom)
			if err != nil {
				return err
			}
			n.appendChild(a)
		}

		n, nodes = md.closeMarker(n, nodes, mark)
	}

	return nil
}
