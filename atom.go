package mobiledoc

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Atom renders an Atom
type Atom func(value string, payload interface{}) string

type atom struct {
	name    string
	value   string
	payload interface{}
}

// UnmarshalJSON decodes the Atom JSON
func (a *atom) UnmarshalJSON(b []byte) error {
	var tmp []json.RawMessage
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return fmt.Errorf("unable to unmarshal atom: %w", err)
	}

	if len(tmp) != 3 {
		return errors.New("atom too short")
	}

	err = json.Unmarshal(tmp[0], &a.name)
	if err != nil {
		return fmt.Errorf("unable to unmarshal atom: %w", err)
	}

	err = json.Unmarshal(tmp[1], &a.value)
	if err != nil {
		return fmt.Errorf("unable to unmarshal atom: %w", err)
	}

	err = json.Unmarshal(tmp[2], &a.payload)
	if err != nil {
		return fmt.Errorf("unable to unmarshal atom: %w", err)
	}

	return nil
}

func (md *Mobiledoc) renderAtom(a *atom) (*node, error) {
	if md.atoms == nil {
		return nil, fmt.Errorf("unable to locate renderer for atom %q", a.name)
	}

	renderer, ok := md.atoms[a.name]
	if !ok {
		return nil, fmt.Errorf("unable to locate renderer for atom %q", a.name)
	}

	return newNode("", renderer(a.value, a.payload)), nil
}
