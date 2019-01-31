// font2go reads one or more standard SVG webfont file(s) and writes Go file(s)
// used to render them to polygons.
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

const (
	prefix = "fonts"
)

var (
	outTemp = template.Must(template.New("out").Funcs(funcMap).Parse(goTemplate))
	funcMap = template.FuncMap{
		"floats":  floats,
		"orEmpty": orEmpty,
		"utf8":    utf8Escape,
	}

	digitRE = regexp.MustCompile(`^\d`)
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Printf("Processing file %q ...", arg)

		fontData := &FontData{}
		if buf, err := ioutil.ReadFile(arg); err != nil {
			log.Fatal(err)
		} else {
			if err := xml.Unmarshal(buf, fontData); err != nil {
				log.Fatal(err)
			}
		}

		fontData.Font.ID = strings.ToLower(fontData.Font.ID)
		fontData.Font.ID = strings.Replace(fontData.Font.ID, "-", "_", -1)
		if digitRE.MatchString(fontData.Font.ID) {
			fontData.Font.ID = "f" + fontData.Font.ID
		}

		sort.Slice(fontData.Font.Glyphs, func(a, b int) bool {
			sa, sb := "", ""
			if fontData.Font.Glyphs[a].Unicode != nil {
				sa = *fontData.Font.Glyphs[a].Unicode
			}
			if fontData.Font.Glyphs[b].Unicode != nil {
				sb = *fontData.Font.Glyphs[b].Unicode
			}
			return strings.Compare(sa, sb) < 0
		})

		for _, g := range fontData.Font.Glyphs {
			g.ParsePath()
			g.GenGerberLP(fontData.Font.FontFace)
		}

		var buf bytes.Buffer
		if err := outTemp.Execute(&buf, fontData.Font); err != nil {
			log.Fatal(err)
		}

		fontDir := filepath.Join(prefix, fontData.Font.ID)
		if err := os.MkdirAll(fontDir, 0755); err != nil {
			log.Fatal(err)
		}
		filename := filepath.Join(fontDir, "font.go")
		fmtBuf, err := format.Source(buf.Bytes())
		if err != nil {
			ioutil.WriteFile(filename, buf.Bytes(), 0644) // Dump the unformatted output.
			log.Fatalf("error formating generated Go code: %v", err)
		}

		if err := ioutil.WriteFile(filename, fmtBuf, 0644); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done.")
}

func utf8Escape(s *string) string {
	if s == nil || *s == "" {
		return `''`
	}
	switch *s {
	case `\`:
		return `'\\'`
	case `'`:
		return `'\''`
	case "fb":
		return `'\ufb1c'` // can't find mapping - made this up.
	case "ff":
		return `'\ufb00'`
	case "ffb":
		return `'\ufb07'` // can't find mapping - made this up.
	case "ffh":
		return `'\ufb08'` // can't find mapping - made this up.
	case "ffi":
		return `'\ufb03'`
	case "ffj":
		return `'\ufb09'` // can't find mapping - made this up.
	case "ffk":
		return `'\ufb0a'` // can't find mapping - made this up.
	case "ffl":
		return `'\ufb04'`
	case "fh":
		return `'\ufb0b'` // can't find mapping - made this up.
	case "fi":
		return `'\ufb01'`
	case "fj":
		return `'\ufb0c'` // can't find mapping - made this up.
	case "fk":
		return `'\ufb0d'` // can't find mapping - made this up.
	case "fl":
		return `'\ufb02'`
	case "ft":
		return `'\ufb05'`
	case "tt":
		return `'\ufb0f'` // can't find mapping - made this up.
	case "1!":
		return `'\ufb10'` // can't find mapping - made this up.
	case "1#":
		return `'\ufb11'` // can't find mapping - made this up.
	case "1$":
		return `'\ufb12'` // can't find mapping - made this up.
	case "1%":
		return `'\ufb13'` // can't find mapping - made this up.
	case "1&":
		return `'\ufb14'` // can't find mapping - made this up.
	case "1(":
		return `'\ufb15'` // can't find mapping - made this up.
	case "1)":
		return `'\ufb16'` // can't find mapping - made this up.
	case "1*":
		return `'\ufb17'` // can't find mapping - made this up.
	case "1@":
		return `'\ufb18'` // can't find mapping - made this up.
	case "1^":
		return `'\ufb19'` // can't find mapping - made this up.
	case "qf":
		return `'\ufb1a'` // can't find mapping - made this up.
	case "qj":
		return `'\ufb1b'` // can't find mapping - made this up.
	case "\ue001\ue014":
		return `'\u2469'`
	case "\ue001\ue015":
		return `'\u246a'`
	case "\ue001\ue016":
		return `'\u246b'`
	case "\ue001\ue017":
		return `'\u246c'`
	case "\ue001\ue018":
		return `'\u246d'`
	case "\ue001\ue019":
		return `'\u246e'`
	case "\ue001\ue01a":
		return `'\u246f'`
	case "\ue001\ue01b":
		return `'\u2470'`
	case "\ue001\ue01c":
		return `'\u2471'`
	case "\ue001\ue01d":
		return `'\u2472'`
	case "\ue002\ue014":
		return `'\u2473'`
	default:
		for _, r := range *s { // Return the first rune
			return fmt.Sprintf("'%c'", r)
		}
	}
	return ""
}

func orEmpty(s *string) string {
	if s == nil || *s == "" {
		return `""`
	}
	return fmt.Sprintf("%q", *s)
}

func floats(f []float64) string {
	return fmt.Sprintf("%#v", f)
}

var goTemplate = `// Auto-generated - DO NOT EDIT!

package {{ .ID }}

import (
	"github.com/gmlewis/go-fonts/fonts"
)

func init() {
  fonts.Fonts["{{ .ID }}"] = {{ .ID }}Font
}

var {{ .ID }}Font = &fonts.Font{
	ID: "{{ .ID }}",
	HorizAdvX:  {{ .HorizAdvX }},
	UnitsPerEm: {{ .FontFace.UnitsPerEm }},
	Ascent:     {{ .FontFace.Ascent }},
	Descent:    {{ .FontFace.Descent }},
	MissingHorizAdvX: {{ .MissingGlyph.HorizAdvX }},
	Glyphs: map[rune]*fonts.Glyph{ {{ range .Glyphs }}{{ if .Unicode }}{{ if .PathSteps }}
		{{ .Unicode | utf8 }}: {
			HorizAdvX: {{ .HorizAdvX }},
			Unicode: {{ .Unicode | utf8 }},
			GerberLP: {{ .GerberLP | orEmpty }},
			PathSteps: []*fonts.PathStep{ {{ range .PathSteps }}
				{ C: '{{ .C }}'{{ if .P }}, P: {{ .P | floats }}{{ end }} },{{ end }}
			},
		},{{ end }}{{ end }}{{ end }}
	},
}
`
