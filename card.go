package mobiledoc

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// CardRenderer renders a Card to the registered format
type CardRenderer func(payload interface{}) string

var cards map[string]map[string]CardRenderer
var cardsLock sync.Mutex

func init() {
	RegisterCardRenderer(
		"image-card",
		"markdown",
		func(payload interface{}) string {
			m, ok := payload.(map[string]interface{})
			if !ok {
				return ""
			}
			src, ok := m["src"]
			if !ok {
				return ""
			}
			return fmt.Sprintf("![](%s)", src.(string))
		},
	)
}

// RegisterCardRenderer registers a CardRenderer for a specific render type
//
// NOTE: currently only "markdown" is accepted as renderType
func RegisterCardRenderer(name, renderType string, renderer CardRenderer) {
	if renderType != "markdown" {
		panic(fmt.Sprintf("unsupported render type %s", renderType))
	}

	cardsLock.Lock()
	defer cardsLock.Unlock()

	if cards == nil {
		cards = make(map[string]map[string]CardRenderer)
	}

	if _, ok := cards[renderType]; !ok {
		cards[renderType] = make(map[string]CardRenderer)
	}

	cards[renderType][name] = renderer
}

type card struct {
	name    string
	payload interface{}
}

// UnmarshalJSON decodes the Card and stores in *c
func (c *card) UnmarshalJSON(b []byte) error {
	var tmp []json.RawMessage
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal card")
	}

	if len(tmp) != 2 {
		return errors.New("card too short")
	}

	err = json.Unmarshal(tmp[0], &c.name)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal card name")
	}

	err = json.Unmarshal(tmp[1], &c.payload)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal card value")
	}

	return nil
}

// Render the card to the specified format
func (c *card) Render(renderType string) (*node, error) {
	if cards == nil {
		return nil, fmt.Errorf("unable to locate renderer for card %q", c.name)
	}

	if _, ok := cards[renderType]; !ok {
		return nil, fmt.Errorf("unable to locate renderer for card %q", c.name)
	}

	if _, ok := cards[renderType][c.name]; !ok {
		return nil, fmt.Errorf("unable to locate renderer for card %q", c.name)
	}

	wrapper := newNode("div", "")
	render := newNode("", cards[renderType][c.name](c.payload))
	wrapper.addChild(render)

	return wrapper, nil
}
