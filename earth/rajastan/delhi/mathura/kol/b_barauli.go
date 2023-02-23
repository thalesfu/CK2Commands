package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗利BarauliBarony struct {
	feud.BaseBarony
}

var BBarauli婆罗利 feud.Barony = &婆罗利BarauliBarony{}

func init() {
	f := BBarauli婆罗利.(*婆罗利BarauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barauli",
		TitleName: "婆罗利",
		TitleCode: "b_barauli",
	}
}
