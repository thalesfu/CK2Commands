package peresechen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔海伊OrheiBarony struct {
	feud.BaseBarony
}

var BOrhei奥尔海伊 feud.Barony = &奥尔海伊OrheiBarony{}

func init() {
	f := BOrhei奥尔海伊.(*奥尔海伊OrheiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orhei",
		TitleName: "奥尔海伊",
		TitleCode: "b_orhei",
	}
}
