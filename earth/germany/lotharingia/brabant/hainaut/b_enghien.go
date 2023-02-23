package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂吉安EnghienBarony struct {
	feud.BaseBarony
}

var BEnghien昂吉安 feud.Barony = &昂吉安EnghienBarony{}

func init() {
	f := BEnghien昂吉安.(*昂吉安EnghienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enghien",
		TitleName: "昂吉安",
		TitleCode: "b_enghien",
	}
}
