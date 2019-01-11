package mobiledoc

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Card renders a Card
type Card func(payload interface{}) string

func imagecard(payload interface{}) string {
	m, ok := payload.(map[string]interface{})
	if !ok {
		return ""
	}
	src, ok := m["src"]
	if !ok {
		return ""
	}
	return fmt.Sprintf("![](%s)", src.(string))
}

type card struct {
	name    string
	payload interface{}
}

// UnmarshalJSON decodes the Card JSON
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
		return errors.Wrap(err, "unable to unmarshal card payload")
	}

	return nil
}

// Render the card to the specified format
func (md *Mobiledoc) renderCard(c *card) (*node, error) {
	if md.cards == nil {
		return nil, fmt.Errorf("unable to locate renderer for card %q", c.name)
	}

	renderer, ok := md.cards[c.name]
	if !ok {
		return nil, fmt.Errorf("unable to locate renderer for card %q", c.name)
	}

	wrapper := newNode("div", "")
	render := newNode("", renderer(c.payload))
	wrapper.addChild(render)

	return wrapper, nil
}
