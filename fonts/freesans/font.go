// Auto-generated - DO NOT EDIT!

package freesans

import (
	"github.com/gmlewis/go-fonts/fonts"
)

var freesansFont = &fonts.Font{
	ID:               "freesans",
	HorizAdvX:        0,
	UnitsPerEm:       1000,
	Ascent:           800,
	Descent:          -200,
	MissingHorizAdvX: 800,
	Glyphs:           map[rune]*fonts.Glyph{},
}

func init() {
	fonts.Fonts["freesans"] = freesansFont
	fonts.InitFromFontData(freesansFont, fontData)
}
