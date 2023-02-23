package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰塔内拉托FontanellatoBarony struct {
	feud.BaseBarony
}

var BFontanellato丰塔内拉托 feud.Barony = &丰塔内拉托FontanellatoBarony{}

func init() {
	f := BFontanellato丰塔内拉托.(*丰塔内拉托FontanellatoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fontanellato",
		TitleName: "丰塔内拉托",
		TitleCode: "b_fontanellato",
	}
}
