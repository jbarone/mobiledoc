package mobiledoc

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var updateFlag bool

func init() {
	flag.BoolVar(&updateFlag, "update", false, "set the update flag to update the expected output of any golden file tests")
}

func TestRenderMarkdown(t *testing.T) {
	RegsiterAtomRenderer(
		"hello-atom",
		"markdown",
		func(value string, payload interface{}) string {
			return fmt.Sprintf("Hello %s", value)
		},
	)
	tests := []string{
		"empty_0.3.0",
		"empty_0.3.1",
		"image_section_0.3.0",
		"image_section_0.3.1",
		"without_markup_0.3.0",
		"without_markup_0.3.1",
		"simple_markup_0.3.0",
		"simple_markup_0.3.1",
		"attribute_markup_0.3.0",
		"attribute_markup_0.3.1",
		"multi_marker_section_0.3.1",
		"list_section_0.3.1",
		"atom_0.3.1",
		"image_card_0.3.1",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			w := &bytes.Buffer{}
			r, err := os.Open(filepath.Join("testdata", tt+".json"))
			if err != nil {
				t.Fatal(err)
			}
			if err = RenderMarkdown(r, w); err != nil {
				t.Errorf("RenderMarkdown() error = %v, want nil", err)
				return
			}

			wantFile := filepath.Join("testdata", "markdown", tt+".golden")
			if updateFlag {
				var f *os.File
				f, err = os.Create(wantFile)
				if err != nil {
					t.Fatalf("os.Create() err = %s; want nil", err)
				}
				f.Write(w.Bytes())
				f.Close()
			}

			f, err := os.Open(wantFile)
			if err != nil {
				t.Fatalf("os.Open() err = %s; want nil", err)
			}
			want, err := ioutil.ReadAll(f)
			f.Close()
			if err != nil {
				t.Fatalf("ioutil.ReadAll() err = %s; want nil", err)
			}

			if gotW := w.Bytes(); !bytes.Equal(gotW, want) {
				t.Errorf("RenderMarkdown() = %q, want %q", gotW, want)
			}
		})
	}
}

func TestRenderMarkdown_errors(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
	}{
		{
			"missing_atom",
			strings.NewReader(`
				{
					"version": "0.3.1",
					"atoms": [
						["missing-atom", "Bob", { "id": 42 }]
					],
					"cards": [],
					"markups": [],
					"sections": [
						[1, "P", [
								[1, [], 0, 0]
							]
						]
					]
				}
			`),
		},
		{
			"missing_card",
			strings.NewReader(`
				{
					"version": "0.3.1",
					"atoms": [],
					"cards": [
						[
							"missing-card",
							{
								"src": "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="
							}
						]
					],
					"markups": [],
					"sections": [
						[10, 0]
					]
				}
			`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := RenderMarkdown(tt.r, w); err == nil {
				t.Errorf("RenderMarkdown() error = %v, wantErr true", err)
			}
		})
	}
}
