package mobiledoc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Marker constants
const (
	markerMarkup = iota
	markerAtom
)

// Section constants
const (
	sectionMarkup = 1
	sectionImage  = 2
	sectionList   = 3
	sectionCard   = 10
)

// Mobiledoc models the data required to render a mobiledoc document
type Mobiledoc struct {
	r     io.Reader
	atoms map[string]Atom
	cards map[string]Card
	doc   doc
	root  *node
}

// NewMobiledoc creates a new Mobiledoc instance
func NewMobiledoc(src io.Reader) Mobiledoc {
	return Mobiledoc{
		r: src,
		cards: map[string]Card{
			"image-card": imagecard,
		},
	}
}

// WithAtom creates a new Mobiledoc instance that has a registered Atom
func (md Mobiledoc) WithAtom(name string, atom Atom) Mobiledoc {
	if md.atoms == nil {
		md.atoms = make(map[string]Atom)
	}
	md.atoms[name] = atom
	return md
}

// WithCard creates a new Mobiledoc instance that has a registered Card
func (md Mobiledoc) WithCard(name string, card Card) Mobiledoc {
	if md.cards == nil {
		md.cards = make(map[string]Card)
	}
	md.cards[name] = card
	return md
}

// Render the Mobiledoc is rendered to the given writer
func (md *Mobiledoc) Render(w io.Writer) error {
	if md.root != nil {
		return md.root.renderMarkdown(w)
	}

	var mdmap map[string]json.RawMessage
	decoder := json.NewDecoder(md.r)
	err := decoder.Decode(&mdmap)
	if err != nil {
		return fmt.Errorf("unable to decode mobiledoc json: %w", err)
	}

	verInt, ok := mdmap["version"]
	if !ok {
		return errors.New("not valid mobiledoc: version not found")
	}

	var version string
	err = json.Unmarshal(verInt, &version)
	if err != nil {
		return fmt.Errorf("not valid mobiledoc: version string: %w", err)
	}

	switch version {
	case "0.3.0", "0.3.1", "0.3.2":
		n, err := md.parseV03(mdmap)
		if err != nil {
			return fmt.Errorf("unable to parse mobiledoc: %w", err)
		}
		md.root = n
	default:
		return fmt.Errorf("unknown version %s", version)
	}
	return md.root.renderMarkdown(w)
}
