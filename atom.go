package mobiledoc

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// AtomRenderer renders an Atom to the registered format
type AtomRenderer func(value string, payload interface{}) string

var atoms map[string]map[string]AtomRenderer
var atomsLock sync.Mutex

// RegsiterAtomRenderer registers an AtomRenderer for a specific render type
//
// NOTE: currently only "markdown" is accepted as renderType
func RegsiterAtomRenderer(name, renderType string, renderer AtomRenderer) {
	if renderType != "markdown" {
		panic(fmt.Sprintf("unsupported render type %s", renderType))
	}

	atomsLock.Lock()
	defer atomsLock.Unlock()
	if atoms == nil {
		atoms = make(map[string]map[string]AtomRenderer)
	}

	if _, ok := atoms[renderType]; !ok {
		atoms[renderType] = make(map[string]AtomRenderer)
	}

	atoms[renderType][name] = renderer
}

type atom struct {
	name    string
	value   string
	payload interface{}
}

// UnmarshalJSON decodes the Atom and stores in *a
func (a *atom) UnmarshalJSON(b []byte) error {
	var tmp []json.RawMessage
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal atom")
	}

	if len(tmp) != 3 {
		return errors.New("atom too short")
	}

	err = json.Unmarshal(tmp[0], &a.name)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal atom name")
	}

	err = json.Unmarshal(tmp[1], &a.value)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal atom value")
	}

	err = json.Unmarshal(tmp[2], &a.payload)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal atom value")
	}

	return nil
}

// Render the atom to the specified format
func (a *atom) Render(renderType string) (*node, error) {
	if atoms == nil {
		return nil, fmt.Errorf("unable to locate renderer for atom %q", a.name)
	}

	if _, ok := atoms[renderType]; !ok {
		return nil, fmt.Errorf("unable to locate renderer for atom %q", a.name)
	}

	if _, ok := atoms[renderType][a.name]; !ok {
		return nil, fmt.Errorf("unable to locate renderer for atom %q", a.name)
	}

	return newNode("", atoms[renderType][a.name](a.value, a.payload)), nil
}
