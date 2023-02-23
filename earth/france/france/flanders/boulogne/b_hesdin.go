package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃丹HesdinBarony struct {
	feud.BaseBarony
}

var BHesdin埃丹 feud.Barony = &埃丹HesdinBarony{}

func init() {
	f := BHesdin埃丹.(*埃丹HesdinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hesdin",
		TitleName: "埃丹",
		TitleCode: "b_hesdin",
	}
}
