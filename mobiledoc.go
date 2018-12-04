package mobiledoc

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
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

// RenderMarkdown renders the given Mobiledoc JSON and renders is as Markdown
// into the given writer
func RenderMarkdown(r io.Reader, w io.Writer) error {
	var mdmap map[string]json.RawMessage
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&mdmap)
	if err != nil {
		return errors.Wrap(err, "unable to decode mobiledoc json")
	}

	verInt, ok := mdmap["version"]
	if !ok {
		return errors.New("not valid mobiledoc: version not found")
	}

	var version string
	err = json.Unmarshal(verInt, &version)
	if err != nil {
		return errors.Wrap(err, "not valid mobiledoc: version string")
	}

	switch version {
	case "0.3.0", "0.3.1":
		n, err := parseV03(mdmap)
		if err != nil {
			return errors.Wrap(err, "unable to parse mobiledoc")
		}
		return n.renderMarkdown(w)
	default:
		return fmt.Errorf("unknown version %s", version)
	}
}
