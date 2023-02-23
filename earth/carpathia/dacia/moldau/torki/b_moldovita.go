package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩尔多维察MoldovitaBarony struct {
	feud.BaseBarony
}

var BMoldovita摩尔多维察 feud.Barony = &摩尔多维察MoldovitaBarony{}

func init() {
	f := BMoldovita摩尔多维察.(*摩尔多维察MoldovitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moldovita",
		TitleName: "摩尔多维察",
		TitleCode: "b_moldovita",
	}
}
