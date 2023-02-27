package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏尔特SyrteBarony struct {
	feud.BaseBarony
}

var BSyrte苏尔特 feud.Barony = &苏尔特SyrteBarony{}

func init() {
    f := BSyrte苏尔特.(*苏尔特SyrteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrte",
		TitleName: "苏尔特",
		TitleCode: "b_syrte",
	}
}
